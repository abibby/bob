package relationships

import (
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestHasMany_has_correct_internal_keys(t *testing.T) {

	f := &Foo{}

	InitializeRelationships(f)

	assert.Equal(t, "id", f.Bars.getParentKey())
	assert.Equal(t, "foo_id", f.Bars.getRelatedKey())
}

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

		err := InitializeRelationships(foos)

		err = Load(tx, foos, "Bars")
		assert.NoError(t, err)

		for _, foo := range foos {
			assert.True(t, foo.Bars.loaded)
			assert.Equal(t,
				[]*Bar{
					{ID: foo.ID * 2, FooID: foo.ID},
					{ID: foo.ID*2 + 1, FooID: foo.ID},
				},
				foo.Bars.value,
			)
		}
	})
}
