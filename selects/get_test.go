package selects_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/bobtesting"
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
		const insert = "INSERT INTO foos (id, name) values (?,?)"
		_, err := tx.ExecContext(context.Background(), insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.ExecContext(context.Background(), insert, 2, "test2")
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
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
		const insert = "INSERT INTO foos (id, name) values (?,?)"
		_, err := tx.ExecContext(context.Background(), insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.ExecContext(context.Background(), insert, 2, "test2")
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
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
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
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
		foos, err := NewTestBuilder().Get(tx)
		assert.NoError(t, err)
		assertJsonEqual(t, `[]`, foos)
	})
}

func TestFirst_returns_nil_with_no_results(t *testing.T) {
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
		foo, err := NewTestBuilder().First(tx)
		assert.NoError(t, err)
		assert.Nil(t, foo)
	})
}
