package selects

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/jmoiron/sqlx"
)

type Relationship interface {
	Subquery() *SubBuilder
	Initialize(self any, field reflect.StructField) error
	Load(ctx context.Context, tx *sqlx.Tx, relations []Relationship) error
	Loaded() bool
}

type relationValue[T any] struct {
	loaded bool
	value  T
}

func (v *relationValue[T]) Value() (T, bool) {
	return v.value, v.loaded
}
func (v *relationValue[T]) Loaded() bool {
	return v.loaded
}

func (v *relationValue[T]) MarshalJSON() ([]byte, error) {
	if !v.loaded {
		return json.Marshal(nil)
	}
	return json.Marshal(v.value)
}

func foreignKeyName(field reflect.StructField, tag string, tableType any) (string, error) {
	t, ok := field.Tag.Lookup(tag)
	if ok {
		return t, nil
	}

	pKeys := builder.PrimaryKey(tableType)
	if len(pKeys) != 1 {
		return "", fmt.Errorf("you must specify keys for relationships with compound primary keys")
	}
	return builder.GetTableSingular(tableType) + "_" + pKeys[0], nil
}

func primaryKeyName(field reflect.StructField, tag string, tableType any) (string, error) {
	t, ok := field.Tag.Lookup(tag)
	if ok {
		return t, nil
	}
	pKeys := builder.PrimaryKey(tableType)
	if len(pKeys) != 1 {
		return "", fmt.Errorf("you must specify keys for relationships with compound primary keys")
	}
	return pKeys[0], nil
}

func getRelation(v any, relation string) (Relationship, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsZero() {
			rv = reflect.New(rv.Type().Elem())
			err := InitializeRelationships(rv.Interface())
			if err != nil {
				panic(err)
			}
		}
		rv = rv.Elem()
	}
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		ft := t.Field(i)

		if ft.Anonymous {
			r, ok := getRelation(rv.Field(i).Interface(), relation)
			if ok {
				return r, true
			}
			continue
		}

		if ft.Name != relation {
			continue
		}

		r, ok := rv.Field(i).Interface().(Relationship)
		if !ok {
			continue
		}
		return r, true
	}
	return nil, false
}
