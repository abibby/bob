package relationships

import (
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects/mysql"
	"github.com/abibby/bob/selects"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
)

type Relationship interface {
	Query() *selects.Builder
	Initialize(self any, field reflect.StructField) error
	Load(tx *sqlx.Tx, relations []Relationship) error
}

type HasOne[T any] struct {
	parent     any
	foreignKey string
	localKey   string
	value      T
}

var _ Relationship = &HasOne[any]{}

func (r *HasOne[T]) Value() (T, error) {
	var zero T

	return zero, nil
}
func (r *HasOne[T]) Query() *selects.Builder {
	var related T
	local, ok := getValue(r.parent, r.localKey)
	if !ok {
		panic(fmt.Errorf("no local key %s", r.localKey))
	}
	return selects.New().
		Select("*").
		From(builder.GetTable(related)).
		Where(r.foreignKey, "=", local).
		Limit(1)
}
func (r *HasOne[T]) Initialize(parent any, field reflect.StructField) error {
	r.parent = parent
	r.localKey, _ = field.Tag.Lookup("local")
	r.foreignKey, _ = field.Tag.Lookup("foreign")
	return nil
}
func (r *HasOne[T]) Load(tx *sqlx.Tx, relations []Relationship) error {
	var related T
	hasOnes := ofType[*HasOne[T]](relations)
	localKeys := make([]any, 0, len(hasOnes))
	for _, r := range hasOnes {
		local, ok := getValue(r.parent, r.localKey)
		if !ok {
			continue
		}
		localKeys = append(localKeys, local)
	}

	relatedLists := []T{}

	err := selects.New().
		Select("*").
		From(builder.GetTable(related)).
		WhereIn(r.foreignKey, localKeys).
		Dump(&mysql.MySQL{}).
		Get(tx, &relatedLists)
	if err != nil {
		return err
	}

	if len(relatedLists) == 0 {
		return nil
	}

	// TODO: replace with something more efficient
	for _, relation := range hasOnes {
		for _, related := range relatedLists {
			local, localOk := getValue(relation.parent, r.localKey)
			foreign, foreignOk := getValue(related, r.foreignKey)
			if localOk && foreignOk && local == foreign {
				relation.value = related
			}

		}
	}

	return nil
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
