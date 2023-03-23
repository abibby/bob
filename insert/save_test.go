package insert

import (
	"testing"

	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
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
