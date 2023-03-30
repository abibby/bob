package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

func main() {
	goSrc, err := ReadSource(".")
	if err != nil {
		log.Fatal(err)
	}

	t := reflect.TypeOf(selects.Builder[models.Model]{})
	src := "package selects\n\n"
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		funsSrc, err := GenerateFuncs(f, goSrc)
		if err != nil {
			log.Fatal(err)
		}
		src += funsSrc
	}

	err = os.WriteFile("generated.go", []byte(src), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadSource(root string) (string, error) {
	dir, err := os.ReadDir(root)
	if err != nil {
		return "", err
	}
	src := ""
	for _, f := range dir {
		if f.Name() == "generated.go" {
			continue
		}
		b, err := os.ReadFile(path.Join(root, f.Name()))
		if err != nil {
			return "", err
		}
		src += string(b)
	}
	return src, nil
}

func GenerateFuncs(f reflect.StructField, goSrc string) (string, error) {
	t := f.Type

	src := ""

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if m.Name == "Clone" || m.Name == "ToSQL" {
			continue
		}

		matches := regexp.MustCompile(fmt.Sprintf(`func \([^)]*\) %s\(([^)]*)\)`, m.Name)).FindStringSubmatch(goSrc)
		if matches == nil {
			return "", fmt.Errorf("could not find %s", m.Name)
		}
		mName := m.Name

		if f.Name == "havings" {
			switch mName {
			case "And":
				mName = "HavingAnd"
			case "Or":
				mName = "HavingOr"
			default:
				mName = strings.ReplaceAll(mName, "Where", "Having")
			}
		}

		params := matches[1]
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
			"func (b *Builder[T]) %s(%s) *Builder[T] {\n"+
				"\tb = b.Clone()\n"+
				"\tb.%s = b.%s.%s(%s)\n"+
				"\treturn b\n"+
				"}\n",
			mName,
			params,
			f.Name,
			f.Name,
			m.Name,
			args,
		)
	}
	return src, nil
}
