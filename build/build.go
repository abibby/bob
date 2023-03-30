package main

import (
	"fmt"
	"reflect"

	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

func main() {
	t := reflect.TypeOf(selects.Builder[models.Model]{})
	src := ""
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		src += GenerateFuncs(f)
	}

	fmt.Print(src)
}

func GenerateFuncs(f reflect.StructField) string {
	t := f.Type

	src := ""

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if m.Name == "Clone" || m.Name == "ToSQL" {
			continue
		}

		params := ""
		args := ""
		for i := 1; i < m.Type.NumIn(); i++ {
			in := m.Type.In(i)
			if i != 1 {
				params += ", "
				args += ", "
			}
			params += fmt.Sprintf("v%d %s", i, in.String())
			args += fmt.Sprintf("v%d", i)
		}

		src += fmt.Sprintf("func (b *Builder[T]) %s(%s) *Builder[T] {\n\tb.%s = b.%s.%s(%s)\n\treturn b\n}\n",
			m.Name,
			params,
			f.Name,
			f.Name,
			m.Name,
			args,
		)
	}
	return src
}
