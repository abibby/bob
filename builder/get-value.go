package builder

import (
	"reflect"
	"strings"
)

func GetValue(v any, key string) (any, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, false
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		if FieldName(rt.Field(i)) == key {
			return rv.Field(i).Interface(), true
		}
	}
	return nil, false
}

func FieldName(f reflect.StructField) string {
	return DBTag(f)[0]
}
func DBTag(f reflect.StructField) []string {
	dbTag, ok := f.Tag.Lookup("db")
	if ok {
		return strings.Split(dbTag, ",")
	}
	return []string{f.Name}
}
