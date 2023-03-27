package selects

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/jmoiron/sqlx"
)

type Relationship interface {
	// Query() *selects.Builder
	Initialize(self any, field reflect.StructField) error
	Load(tx *sqlx.Tx, relations []Relationship) error
	Loaded() bool
}

type relationValue[T any] struct {
	loaded bool
	value  T
}

func (v *relationValue[T]) Value(tx *sqlx.Tx, r Relationship) (T, error) {
	if !v.loaded {
		err := r.Load(tx, []Relationship{r})
		if err != nil {
			var zero T
			return zero, err
		}
	}
	return v.value, nil
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
