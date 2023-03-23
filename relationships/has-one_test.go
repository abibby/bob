package relationships

import (
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestHasOne_has_correct_internal_keys(t *testing.T) {
	f := &Foo{}

	InitializeRelationships(f)

	assert.Equal(t, "foo_id", f.Bar.getRelatedKey())
	assert.Equal(t, "id", f.Bar.getParentKey())
}

// func TestHasOne(t *testing.T) {
// 	foo := &Foo{
// 		ID: 10,
// 	}
// 	err := InitializeRelationships(foo)
// 	assert.NoError(t, err)

// 	test.QueryTest(t, []test.Case{
// 		{
// 			Name:             "has one",
// 			Builder:          foo.Bar.Query(),
// 			ExpectedSQL:      "SELECT * FROM `Bar` WHERE `foo_id` = ? LIMIT 1",
// 			ExpectedBindings: []any{foo.ID},
// 		},
// 	})
// }

func TestHasOneLoad(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		foos := []*Foo{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}
		for _, f := range foos {
			insert.Save(tx, f)
		}
		insert.Save(tx, &test.Bar{ID: 4, FooID: 1})
		insert.Save(tx, &test.Bar{ID: 5, FooID: 2})
		insert.Save(tx, &test.Bar{ID: 6, FooID: 3})

		err := InitializeRelationships(foos)

		err = Load(tx, foos, "Bar")
		assert.NoError(t, err)

		for _, foo := range foos {
			assert.Equal(t, &Bar{ID: foo.ID + 3, FooID: foo.ID}, foo.Bar.value)
			assert.True(t, foo.Bar.loaded)
		}
	})
}

// type Foo struct {
// 	ID int `db:"id"`
// }

// func (f *Foo) Bar() *HasOne[*Foo, *Bar] {
// 	return NewHasOne(f, &Bar{}, "foo_id", "id")
// }

// type Bar struct {
// 	FooID int `db:"foo_id"`
// }

// func (f *Bar) Foo() *BelongsTo[*Foo] {
// 	return NewBelongsTo(&Foo{})
// }

// func TestHasOne(t *testing.T) {
// 	foo := &Foo{
// 		ID: 10,
// 	}
// 	test.QueryTest(t, []test.Case{
// 		{
// 			"has one",
// 			foo.Bar(),
// 			"SELECT * FROM `Bar` WHERE `foo_id` = ? LIMIT 1",
// 			[]any{foo.ID},
// 		},
// 	})
// }
