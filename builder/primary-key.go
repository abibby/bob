package builder

import "reflect"

type PrimaryKeyer interface {
	PrimaryKey() []string
}

func PrimaryKey(m any) []string {
	if m, ok := m.(PrimaryKeyer); ok {
		return m.PrimaryKey()
	}

	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fallback := ""
	primary := []string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			continue
		}
		tag := DBTag(f)
		if fallback == "" {
			fallback = tag[0]
		}
		if includes(tag, "primary") {
			primary = append(primary, tag[0])
		}
	}
	if len(primary) == 0 {
		return []string{fallback}
	}

	return primary
}

func includes[T comparable](arr []T, v T) bool {
	for _, e := range arr {
		if v == e {
			return true
		}
	}
	return false
}
