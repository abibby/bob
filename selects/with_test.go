package selects_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type WithFoo struct {
	test.Foo
	WithBar *selects.HasOne[*WithBar] `db:"-"`
}

func (f *WithFoo) Table() string {
	return "foos"
}

type WithBar struct {
	test.Bar
}

func (f *WithBar) Table() string {
	return "bars"
}

func (f *WithBar) Scopes() []*selects.Scope {
	return []*selects.Scope{
		{
			Name: "test",
			Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
				return b.Where("id", "=", b.Context().Value("id"))
			},
		},
	}
}

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

func TestWith_HasOne_anonymous(t *testing.T) {

	test.WithDatabase(func(tx *sqlx.Tx) {
		MustSave(tx, &test.Foo{ID: 1})
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		_, err := selects.From[*WithFoo]().With("Bar").Where("id", "=", 1).Get(tx)

		assert.NoError(t, err)
	})
}

func TestWith_HasOne_scope_context(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		MustSave(tx, &test.Foo{ID: 1})
		MustSave(tx, &test.Bar{ID: 2, FooID: 1})
		MustSave(tx, &test.Bar{ID: 3, FooID: 1})
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		ctx := context.WithValue(context.Background(), "id", 3)
		f, err := selects.From[*WithFoo]().
			With("WithBar").
			WithContext(ctx).
			First(tx)
		if !assert.NoError(t, err) {
			return
		}

		b, ok := f.WithBar.Value()
		assert.True(t, ok)
		assert.NotNil(t, b)
		assert.Equal(t, 3, b.ID)
	})
}
