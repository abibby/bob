package relationships

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
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
		spew.Dump("not struct", rv.Kind())
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
