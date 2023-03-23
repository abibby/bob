package builder

import (
	"reflect"
	"strings"

	strcase "github.com/stoewer/go-strcase"
)

type Tabler interface {
	Table() string
}

func GetTable(m any) string {
	name := GetTableSingular(m)
	if !strings.HasSuffix(name, "s") {
		name += "s"
	}
	return name
}

func GetTableSingular(m any) string {
	if m, ok := m.(Tabler); ok {
		return m.Table()
	}
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strcase.SnakeCase(t.Name())
}
