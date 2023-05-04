package schema_test

import (
	"testing"

	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/test"
)

func TestUpdateTable(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "empty update",
			Builder:          schema.Table("foo", func(table *schema.Blueprint) {}),
			ExpectedSQL:      "",
			ExpectedBindings: []any{},
		},
		{
			Name: "add column",
			Builder: schema.Table("foo", func(table *schema.Blueprint) {
				table.String("bar")
			}),
			ExpectedSQL:      "ALTER TABLE \"foo\" ADD \"bar\" TEXT NOT NULL;",
			ExpectedBindings: []any{},
		},
		{
			Name: "change column",
			Builder: schema.Table("foo", func(table *schema.Blueprint) {
				table.Int("id").Change()
			}),
			ExpectedSQL:      "ALTER TABLE \"foo\" MODIFY COLUMN \"id\" INTEGER NOT NULL;",
			ExpectedBindings: []any{},
		},
	})
}
