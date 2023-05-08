package migrate_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/bobtesting"
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	_ "github.com/abibby/bob/test"
)

func TestMigrations(t *testing.T) {
	bobtesting.RunWithDatabase(t, "dont rerun migraitons", func(t *testing.T, tx *sqlx.Tx) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "1",
			Up: func() builder.ToSQLer {
				return schema.Create("foo", func(b *schema.Blueprint) {
					b.Int("id").Primary()
				})
			},
		})

		err := m.Up(context.Background(), tx)
		assert.NoError(t, err)
		m.Add(&migrate.Migration{
			Name: "2",
			Up: func() builder.ToSQLer {
				return schema.Table("foo", func(b *schema.Blueprint) {
					b.String("name")
				})
			},
		})

		err = m.Up(context.Background(), tx)
		assert.NoError(t, err)
	})
}
