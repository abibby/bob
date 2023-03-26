package selects

import (
	"encoding/json"
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