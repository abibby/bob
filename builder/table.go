package builder

import (
	"reflect"
)

type Tabler interface {
	Table() string
}

func GetTable(m any) string {
	if m, ok := m.(Tabler); ok {
		return m.Table()
	}
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
