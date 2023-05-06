package migrations

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: "20230505_222422-Bar",
		Up: func() builder.ToSQLer {
			return schema.Create("bars", func(table *schema.Blueprint) {
				table.Int("id").Primary().AutoIncrement()
				table.Int("foo_id")
			})
		},
		Down: func() builder.ToSQLer {
			return schema.DropIfExists("bars")
		},
	})
}
