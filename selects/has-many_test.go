package selects_test

import (
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestHasMany_Load(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foos := []*Foo{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}
		for _, f := range foos {
			assert.NoError(t, insert.Save(tx, f))
		}
		bars := []*Bar{
			{ID: 2, FooID: 1},
			{ID: 3, FooID: 1},
			{ID: 4, FooID: 2},
			{ID: 5, FooID: 2},
			{ID: 6, FooID: 3},
			{ID: 7, FooID: 3},
		}
		for _, b := range bars {
			assert.NoError(t, insert.Save(tx, b))
		}

		err := selects.Load(tx, foos, "Bars")
		assert.NoError(t, err)

		for _, foo := range foos {
			assert.True(t, foo.Bars.Loaded())
			bars, err := foo.Bars.Value(nil)
			assert.NoError(t, err)
			assert.Len(t, bars, 2)
			assert.Equal(t, foo.ID*2, bars[0].ID)
			assert.Equal(t, foo.ID, bars[0].FooID)
			assert.Equal(t, foo.ID*2+1, bars[1].ID)
			assert.Equal(t, foo.ID, bars[1].FooID)
		}
	})
}
