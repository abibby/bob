package insert_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/bobtesting"
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/hooks"
	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/stretchr/testify/assert"
)

func TestSave_create(t *testing.T) {
	bobtesting.RunWithDatabase(t, "create", func(t *testing.T, tx builder.QueryExecer) {
		f := &test.Foo{
			ID:   1,
			Name: "test",
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		rows, err := tx.QueryContext(context.Background(), "select id, name from foos")
		assert.NoError(t, err)

		assert.True(t, rows.Next())
		id := 0
		name := ""
		rows.Scan(&id, &name)

		assert.Equal(t, f.ID, id)
		assert.Equal(t, f.Name, name)

		assert.False(t, rows.Next())
	})
}

func TestSave_update(t *testing.T) {
	bobtesting.RunWithDatabase(t, "update", func(t *testing.T, tx builder.QueryExecer) {
		f := &test.Foo{
			ID:   1,
			Name: "test",
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		f.Name = "new name"
		err = insert.Save(tx, f)
		assert.NoError(t, err)

		rows, err := tx.QueryContext(context.Background(), "select id, name from foos")
		assert.NoError(t, err)

		assert.True(t, rows.Next())
		id := 0
		name := ""
		rows.Scan(&id, &name)

		assert.Equal(t, f.ID, id)
		assert.Equal(t, f.Name, name)

		assert.False(t, rows.Next())
	})
}

func TestSave_model_is_in_database_after_saving(t *testing.T) {
	bobtesting.RunWithDatabase(t, "model in database after saving", func(t *testing.T, tx builder.QueryExecer) {
		f := &test.Foo{
			ID: 1,
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.InDatabase())
	})
}

func TestSave_autoincrement(t *testing.T) {
	bobtesting.RunWithDatabase(t, "autoincrement", func(t *testing.T, tx builder.QueryExecer) {
		f := &test.Foo{}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		assert.Equal(t, f.ID, 1)
	})
	bobtesting.RunWithDatabase(t, "autoincrement set id", func(t *testing.T, tx builder.QueryExecer) {
		f := &test.Foo{
			ID: 100,
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		assert.Equal(t, f.ID, 100)
	})
}

type FooSaveHookTest struct {
	test.Foo
	saved bool
}

type FooSaveHookTestWrapper struct {
	FooSaveHookTest
}

var _ hooks.AfterSaver = &FooSaveHookTest{}

func (f *FooSaveHookTest) AfterSave(context.Context, builder.QueryExecer) error {
	f.saved = true
	return nil
}
func (f *FooSaveHookTest) Table() string {
	return "foos"
}

func TestSave_hooks(t *testing.T) {
	bobtesting.RunWithDatabase(t, "runs hooks", func(t *testing.T, tx builder.QueryExecer) {
		f := &FooSaveHookTest{
			Foo: test.Foo{
				ID: 1,
			},
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.saved)
	})

	bobtesting.RunWithDatabase(t, "runs hooks on anonymise structs", func(t *testing.T, tx builder.QueryExecer) {
		f := &FooSaveHookTestWrapper{
			FooSaveHookTest{
				Foo: test.Foo{
					ID: 1,
				},
			},
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.saved)
	})
}

type SaveFooReadonly struct {
	test.Foo
	Readonly string `db:"computed,readonly"`
}

func (f *SaveFooReadonly) Table() string {
	return "foos"
}

func TestSave_readonly(t *testing.T) {
	bobtesting.RunWithDatabase(t, "runs hooks", func(t *testing.T, tx builder.QueryExecer) {
		f := &SaveFooReadonly{
			Foo: test.Foo{
				ID: 1,
			},
			Readonly: "yes",
		}
		err := insert.Save(tx, f)
		assert.NoError(t, err)

		newFoo, err := selects.From[*SaveFooReadonly]().First(tx)
		assert.NoError(t, err)
		assert.Equal(t, "", newFoo.Readonly)
	})

}
