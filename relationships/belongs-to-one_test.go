package relationships

import (
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBelongsTo_has_correct_internal_keys(t *testing.T) {

	b := &Bar{}

	InitializeRelationships(b)

	assert.Equal(t, "foo_id", b.Foo.getParentKey())
	assert.Equal(t, "id", b.Foo.getRelatedKey())
}

func TestBelongsToLoad(t *testing.T) {
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
			{ID: 4, FooID: 1},
			{ID: 5, FooID: 2},
			{ID: 6, FooID: 3},
		}
		for _, b := range bars {
			assert.NoError(t, insert.Save(tx, b))
		}

		err := InitializeRelationships(bars)

		err = Load(tx, bars, "Foo")
		if !assert.NoError(t, err) {
			return
		}

		for _, bar := range bars {
			assert.Equal(t, &Foo{ID: bar.FooID}, bar.Foo.value)
			assert.True(t, bar.Foo.loaded)
		}
	})
}
