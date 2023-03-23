package selects

import (
	"testing"

	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		const insert = "INSERT INTO foo (id, name) values (?,?)"
		_, err := tx.Exec(insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.Exec(insert, 2, "test2")
		assert.NoError(t, err)

		foos := []test.Foo{}
		err = New().Select("*").From("foo").Get(tx, &foos)
		assert.NoError(t, err)
		assert.Equal(t, foos, []test.Foo{
			{ID: 1, Name: "test1"},
			{ID: 2, Name: "test2"},
		})
	})
}

func TestFirst(t *testing.T) {
	test.WithDatabase(func(tx *sqlx.Tx) {
		const insert = "INSERT INTO foo (id, name) values (?,?)"
		_, err := tx.Exec(insert, 1, "test1")
		assert.NoError(t, err)
		_, err = tx.Exec(insert, 2, "test2")
		assert.NoError(t, err)

		foo := &test.Foo{}
		err = New().Select("*").From("foo").First(tx, foo)
		assert.NoError(t, err)
		assert.Equal(t, foo, &test.Foo{ID: 1, Name: "test1"})
	})
}
