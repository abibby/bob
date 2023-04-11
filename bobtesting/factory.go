package bobtesting

import (
	"fmt"
	"reflect"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

var factories = map[string]any{}

type Factory[T models.Model] struct {
	make func() T
}
type CountFactory[T models.Model] struct {
	factory *Factory[T]
	count   int
}

type State[T models.Model] interface {
	Apply(T) T
}
type stateFunc[T models.Model] func(T) T

func (sf stateFunc[T]) Apply(t T) T {
	return sf(t)
}

func StateFunc[T models.Model](apply func(T) T) State[T] {
	return stateFunc[T](apply)
}

func name[T any]() string {
	var m T
	return reflect.TypeOf(m).String()
}

func DefineFactory[T models.Model](cb func() T) {
	factories[name[T]()] = &Factory[T]{
		make: cb,
	}
}
func NewFactory[T models.Model]() *Factory[T] {
	f, ok := factories[name[T]()]
	if !ok {
		panic(fmt.Errorf("No factory found for %s", name[T]()))
	}

	return f.(*Factory[T])
}

func (f *Factory[T]) Count(count int) *CountFactory[T] {
	return &CountFactory[T]{
		factory: f,
		count:   count,
	}
}
func (f *Factory[T]) State(s State[T]) *Factory[T] {
	return &Factory[T]{
		make: func() T {
			return s.Apply(f.make())
		},
	}
}

func (f *Factory[T]) Create(tx *sqlx.Tx) T {
	m := f.make()
	err := insert.Save(tx, m)
	if err != nil {
		panic(err)
	}
	return m
}

func (f *CountFactory[T]) Create(tx *sqlx.Tx) []T {
	models := make([]T, f.count)
	for i := 0; i < f.count; i++ {
		m := f.factory.make()
		err := insert.Save(tx, m)
		if err != nil {
			panic(err)
		}
		models[i] = m
	}
	return models
}
