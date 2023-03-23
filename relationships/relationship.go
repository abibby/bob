package relationships

import (
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/jmoiron/sqlx"
)

type Relationship interface {
	// Query() *selects.Builder
	Initialize(self any, field reflect.StructField) error
	Load(tx *sqlx.Tx, relations []Relationship) error
}

func getValue(v any, key string) (any, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, false
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		dbTag, ok := rt.Field(i).Tag.Lookup("db")
		if ok && dbTag == key {
			return rv.Field(i).Interface(), true
		}
	}
	return nil, false
}

func foreignKeyName(field reflect.StructField, tag string, tableType any) string {
	t, ok := field.Tag.Lookup(tag)
	if ok {
		return t
	}
	return builder.GetTableSingular(tableType) + "_id"
}

func primaryKeyName(field reflect.StructField, tag string) string {
	t, ok := field.Tag.Lookup(tag)
	if ok {
		return t
	}
	return "id"
}
