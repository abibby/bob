package selects_test

import (
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestWith_HasOne(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		MustSave(tx, &test.Foo{ID: 1})
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		f, err := NewTestBuilder().With("Bar").Where("id", "=", 1).Get(tx)

		assert.NoError(t, err)
		assertJsonEqual(t, `[{
			"id":1,
			"name":"",
			"bar":{"id":4,"foo_id":1,"foo":null},
			"bars":null
		}]`, f)

	})
}

func TestWith_HasOne_bad_relation(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		MustSave(tx, &test.Foo{ID: 1})
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		_, err := NewTestBuilder().With("BadRelation").Where("id", "=", 1).Get(tx)

		assert.ErrorIs(t, err, selects.ErrMissingRelationship)
	})
}

type WithTestFoo struct {
	test.Foo
}

func (f *WithTestFoo) Table() string {
	return "foos"
}

func TestWith_HasOne_anonymous(t *testing.T) {

	test.WithDatabase(func(tx *sqlx.Tx) {
		MustSave(tx, &test.Foo{ID: 1})
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		_, err := selects.From[*WithTestFoo]().With("Bar").Where("id", "=", 1).Get(tx)

		assert.NoError(t, err)
	})
}
