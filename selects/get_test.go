package selects_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		const insert = "INSERT INTO foos (id, name) values (?,?)"
		_, err := tx.Exec(insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.Exec(insert, 2, "test2")
		assert.NoError(t, err)

		foos, err := selects.From[*test.Foo]().Get(tx)
		assert.NoError(t, err)
		assertJsonEqual(t, `[
			{"id":1,"name":"test1","bar":null,"bars":null},
			{"id":2,"name":"test2","bar":null,"bars":null}
		]`, foos)
	})
}

func TestFirst(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		const insert = "INSERT INTO foos (id, name) values (?,?)"
		_, err := tx.Exec(insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.Exec(insert, 2, "test2")
		assert.NoError(t, err)

		foo, err := selects.From[*test.Foo]().First(tx)
		assert.NoError(t, err)
		assertJsonEqual(t, `{
			"id":1,
			"name":"test1",
			"bar":null,
			"bars":null
		}`, foo)
	})
}

func TestGet_with_scope_and_context(t *testing.T) {
	scopeCtx := &selects.Scope{
		Name: "ctx",
		Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
			return b.Where("id", "=", b.Context().Value("foo"))
		},
	}
	test.WithDatabase(func(tx *sqlx.Tx) {
		ctx := context.WithValue(context.Background(), "foo", 2)

		MustSave(tx, &test.Foo{ID: 1, Name: "foo1"})
		MustSave(tx, &test.Foo{ID: 2, Name: "foo2"})

		foos, err := NewTestBuilder().
			WithScope(scopeCtx).
			Where("name", "like", "foo%").
			WithContext(ctx).
			Get(tx)
		assert.NoError(t, err)
		assertJsonEqual(t, `[{
			"id":2,
			"name":"foo2",
			"bar":null,
			"bars":null
		}]`, foos)
	})

}

func TestGet_returns_empty_array_with_no_results(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foos, err := NewTestBuilder().Get(tx)
		assert.NoError(t, err)
		assertJsonEqual(t, `[]`, foos)
	})
}

func TestFirst_returns_nil_with_no_results(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foo, err := NewTestBuilder().First(tx)
		assert.NoError(t, err)
		assert.Nil(t, foo)
	})
}
