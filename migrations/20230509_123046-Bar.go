package migrations

import (
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: "20230509_123046-Bar",
		Up: schema.Create("bars", func(table *schema.Blueprint) {
			table.Int("id").Primary().AutoIncrement()
			table.Int("foo_id")
		}),
		Down: schema.DropIfExists("bars"),
	})
}
