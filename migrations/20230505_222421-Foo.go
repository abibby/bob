package migrations

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: "20230505_222421-Foo",
		Up: func() builder.ToSQLer {
			return schema.Create("foos", func(table *schema.Blueprint) {
				table.Int("id").Primary().AutoIncrement()
				table.String("name")
			})
		},
		Down: func() builder.ToSQLer {
			return schema.DropIfExists("foos")
		},
	})
}
