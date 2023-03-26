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
		foos := []*test.Foo{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}
		for _, f := range foos {
			assert.NoError(t, insert.Save(tx, f))
		}
		bars := []*test.Bar{
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

		assertJsonEqual(t, `[
			{
			  "id": 1,
			  "name": "",
			  "bar": null,
			  "bars": [
				{ "id": 2, "foo_id": 1, "foo": null },
				{ "id": 3, "foo_id": 1, "foo": null }
			  ]
			},
			{
			  "id": 2,
			  "name": "",
			  "bar": null,
			  "bars": [
				{ "id": 4, "foo_id": 2, "foo": null },
				{ "id": 5, "foo_id": 2, "foo": null }
			  ]
			},
			{
			  "id": 3,
			  "name": "",
			  "bar": null,
			  "bars": [
				{ "id": 6, "foo_id": 3, "foo": null },
				{ "id": 7, "foo_id": 3, "foo": null }
			  ]
			}
		  ]`, foos)
		for _, foo := range foos {
			assert.True(t, foo.Bars.Loaded())
		}
	})
}