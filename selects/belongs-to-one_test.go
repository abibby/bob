package selects_test

import (
	"testing"

	"github.com/abibby/bob/bobtesting"
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBelongsToLoad(t *testing.T) {
	bobtesting.RunWithDatabase(t, "", func(t *testing.T, tx builder.QueryExecer) {
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
			foo, ok := bar.Foo.Value()
			assert.True(t, ok)
			assert.Equal(t, bar.FooID, foo.ID)
		}
	})

	bobtesting.RunWithDatabase(t, "uuids", func(t *testing.T, tx builder.QueryExecer) {
		type Foo struct {
			models.BaseModel
			ID   int       `json:"id" db:"id,primary,autoincrement"`
			Name uuid.UUID `json:"name" db:"name"`
		}
		type Bar struct {
			models.BaseModel
			FooName *uuid.UUID               `json:"foo_id" db:"foo_id"`
			Foo     *selects.BelongsTo[*Foo] `json:"foo"    db:"-" foreign:"foo_id" owner:"name"`
		}

		bars := []*Bar{
			{FooName: newUUID()},
			{FooName: newUUID()},
			{FooName: nil},
			{FooName: nil},
		}
		for _, b := range bars {
			if b.FooName != nil {
				MustSave(tx, &Foo{Name: *b.FooName})
			}
		}
		selects.InitializeRelationships(bars)
		selects.Load(tx, bars, "Foo")

		for i, b := range bars {
			f, ok := b.Foo.Value()
			assert.True(t, ok)
			if i < 2 {
				assert.NotNil(t, f)
			} else {
				assert.Nil(t, f)
			}
		}
	})
}

func newUUID() *uuid.UUID {
	id := uuid.New()
	return &id
}
