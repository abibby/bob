package bob

import (
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/selects"
)

type Relationship[T any] interface {
	builder.ToSQLer
	Query() *selects.Builder
	Value() (T, error)
}

type HasOne[S, T any] struct {
	self       S
	foreignKey string
	localKey   string
}

func NewHasOne[S, T any](self S, m T, foreignKey, localKey string) *HasOne[S, T] {
	return &HasOne[S, T]{
		self:       self,
		foreignKey: foreignKey,
		localKey:   localKey,
	}
}

func (r *HasOne[S, T]) Value() (T, error) {
	var zero T

	return zero, nil
}
func (r *HasOne[S, T]) ToSQL(d dialects.Dialect) (string, []any, error) {
	var zero T
	local, ok := getValue(r.self, r.localKey)
	if !ok {
		return "", nil, fmt.Errorf("no local key %s", r.localKey)
	}
	return From(zero).
		Where(r.foreignKey, "=", local).
		Limit(1).
		ToSQL(d)
}

type BelongsTo[T any] struct{}

func NewBelongsTo[T any](m T) *BelongsTo[T] {
	return &BelongsTo[T]{}
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
