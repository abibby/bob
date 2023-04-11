package bobtesting_test

import (
	"testing"

	"github.com/abibby/bob/bobtesting"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	bobtesting.DefineFactory(func() *test.Foo {
		return &test.Foo{
			Name: "foo",
		}
	})
	bobtesting.RunWithDatabase(t, "create", func(t *testing.T, tx *sqlx.Tx) {
		f := bobtesting.NewFactory[*test.Foo]().Create(tx)
		assert.Equal(t, "foo", f.Name)

		dbF, err := selects.From[*test.Foo]().Find(tx, f.ID)
		assert.NoError(t, err)
		assert.Equal(t, f, dbF)
	})
	bobtesting.RunWithDatabase(t, "count", func(t *testing.T, tx *sqlx.Tx) {
		foos := bobtesting.NewFactory[*test.Foo]().Count(4).Create(tx)
		assert.Len(t, foos, 4)
		for _, f := range foos {
			assert.Equal(t, "foo", f.Name)
		}
	})
	bobtesting.RunWithDatabase(t, "state", func(t *testing.T, tx *sqlx.Tx) {
		f := bobtesting.NewFactory[*test.Foo]().
			State(bobtesting.StateFunc(func(f *test.Foo) *test.Foo {
				f.Name = "bar"
				return f
			})).
			Create(tx)
		assert.Equal(t, "bar", f.Name)
	})
}
