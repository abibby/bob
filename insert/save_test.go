package insert

import (
	"context"
	"testing"

	"github.com/abibby/bob/hooks"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestSave_create(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &test.Foo{
			ID:   1,
			Name: "test",
		}
		err := Save(tx, f)
		assert.NoError(t, err)

		rows, err := tx.Query("select id, name from foos")
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
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &test.Foo{
			ID:   1,
			Name: "test",
		}
		err := Save(tx, f)
		assert.NoError(t, err)

		f.Name = "new name"
		err = Save(tx, f)
		assert.NoError(t, err)

		rows, err := tx.Query("select id, name from foos")
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
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &test.Foo{
			ID: 1,
		}
		err := Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.InDatabase())
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

func (f *FooSaveHookTest) AfterSave(context.Context, *sqlx.Tx) error {
	f.saved = true
	return nil
}
func (f *FooSaveHookTest) Table() string {
	return "foos"
}

func TestSave_runs_hooks(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &FooSaveHookTest{
			Foo: test.Foo{
				ID: 1,
			},
		}
		err := Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.saved)
	})
}

func TestSave_runs_hooks_on_anonymise_structs(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		f := &FooSaveHookTestWrapper{
			FooSaveHookTest{
				Foo: test.Foo{
					ID: 1,
				},
			},
		}
		err := Save(tx, f)
		assert.NoError(t, err)

		assert.True(t, f.saved)
	})
}
