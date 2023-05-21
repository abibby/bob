package migrations

import (
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: "20230521_094632-Bar",
		Up: schema.Table("bars", func(table *schema.Blueprint) {
			table.ForeignKey("foo_id", "foos", "id")
		}),
		Down: schema.Table("bars", func(table *schema.Blueprint) {
		}),
	})
}
