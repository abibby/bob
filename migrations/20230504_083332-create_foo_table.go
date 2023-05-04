package migrations

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
)

func init() {
	migrate.New(migrate.Migration{
		Name: "20230504_083332-create_foo_table",
		Up: func() builder.ToSQLer {
			return schema.Table("", func(table *schema.Blueprint) {
				table.Bool("test").Change()
			})
		},
		Down: func() builder.ToSQLer {
			return schema.Table("", func(table *schema.Blueprint) {
			})
		},
	})
}
