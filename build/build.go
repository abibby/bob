package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	line, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		panic(err)
	}
	file := os.Getenv("GOFILE")

	structName, structParams, structFields, err := GetStruct(file, line)
	if err != nil {
		panic(err)
	}

	goSrc, err := ReadSource(".")
	if err != nil {
		panic(err)
	}
	matches := regexp.MustCompile(fmt.Sprintf(`func \((\w+ +)?([^)]+)\) ([\w)]+)\((.*)\) \*?\w+(:?\[.+\])? {`)).FindAllStringSubmatch(goSrc, -1)
	src := "package selects\n\n"

	for _, match := range matches {
		fieldType := match[2]
		methodName := match[3]
		params := match[4]

		if methodName == "Clone" || methodName == "ToSQL" || !IsUppercase(methodName[0]) {
			continue
		}
		if fieldNames, ok := structFields[fieldType]; ok {
			for _, fieldName := range fieldNames {
				originalMethodName := methodName
				if fieldName == "havings" {
					switch methodName {
					case "And":
						methodName = "HavingAnd"
					case "Or":
						methodName = "HavingOr"
					default:
						methodName = strings.ReplaceAll(methodName, "Where", "Having")
					}
				}
				args := ""
				for i, p := range strings.Split(params, ",") {
					if i != 0 {
						args += ", "
					}
					parts := strings.Split(strings.TrimSpace(p), " ")
					args += parts[0]
					if len(parts) > 1 && strings.HasPrefix(parts[1], "...") {
						args += "..."
					}
				}
				src += fmt.Sprintf(
					"func (b *%s) %s(%s) *%s {\n"+
						"\tb = b.Clone()\n"+
						"\tb.%s = b.%s.%s(%s)\n"+
						"\treturn b\n"+
						"}\n",
					structName+structParams,
					methodName,
					params,
					structName+structParams,
					fieldName,
					fieldName,
					originalMethodName,
					args,
				)
			}
		}
	}

	err = os.WriteFile(fmt.Sprintf("generated_%s.go", structName), []byte(src), 0644)
	if err != nil {
		panic(err)
	}
}

func IsUppercase(r byte) bool {
	return r >= 'A' && r <= 'Z'
}

func GetStruct(file string, line int) (string, string, map[string][]string, error) {
	m := map[string][]string{}

	b, err := os.ReadFile(file)
	if err != nil {
		return "", "", nil, err
	}
	src := string(b)

	lines := strings.Split(src, "\n")

	matches := regexp.MustCompile(`type (\w*)(\[.*\])? struct {`).FindStringSubmatch(lines[line])
	structName := matches[1]
	params := matches[2]
	if params != "" {
		params = "[T]"
	}

	for _, l := range lines[line+1:] {
		if l == "}" {
			return structName, params, m, nil
		}
		parts := strings.Fields(strings.TrimSpace(l))
		if len(parts) < 2 {
			continue
		}
		a, ok := m[parts[1]]
		if !ok {
			a = make([]string, 0, 1)
		}
		a = append(a, parts[0])
		m[parts[1]] = a
	}

	return "", "", nil, fmt.Errorf("No struct found in %s:%d", file, line)
}

func ReadSource(root string) (string, error) {
	dir, err := os.ReadDir(root)
	if err != nil {
		return "", err
	}
	src := ""
	for _, f := range dir {
		b, err := os.ReadFile(path.Join(root, f.Name()))
		if err != nil {
			return "", err
		}
		src += string(b)
	}
	return src, nil
}
