package selects_test

import (
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestHasOneLoad(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foos := []*test.Foo{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}
		for _, f := range foos {
			MustSave(tx, f)
		}
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})
		MustSave(tx, &test.Bar{ID: 5, FooID: 2})
		MustSave(tx, &test.Bar{ID: 6, FooID: 3})

		err := selects.Load(tx, foos, "Bar")
		assert.NoError(t, err)

		for _, foo := range foos {
			assert.True(t, foo.Bar.Loaded())
			bar, err := foo.Bar.Value(nil)
			assert.NoError(t, err)
			assert.Equal(t, foo.ID+3, bar.ID)
			assert.Equal(t, foo.ID, bar.FooID)
		}
	})
}

func TestHasOne_json_marshal(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &test.Foo{ID: 1}
		MustSave(tx, f)
		MustSave(tx, &test.Bar{ID: 4, FooID: 1})

		err := selects.Load(tx, f, "Bar")
		assert.NoError(t, err)

		assertJsonEqual(t, `{
			"id":1,
			"name":"",
			"bar":{"id":4,"foo_id":1,"foo":null},
			"bars":null
		}`, f)

	})

}
