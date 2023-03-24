package selects_test

import (
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBelongsToLoad(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foos := []*test.Foo{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}
		for _, f := range foos {
			assert.NoError(t, insert.Save(tx, f))
		}
		bars := []*test.Bar{
			{ID: 4, FooID: 1},
			{ID: 5, FooID: 2},
			{ID: 6, FooID: 3},
		}
		for _, b := range bars {
			assert.NoError(t, insert.Save(tx, b))
		}

		err := selects.Load(tx, bars, "Foo")
		if !assert.NoError(t, err) {
			return
		}

		for _, bar := range bars {
			assert.True(t, bar.Foo.Loaded())
			foo, err := bar.Foo.Value(nil)
			assert.NoError(t, err)
			assert.Equal(t, bar.FooID, foo.ID)
		}
	})
}
