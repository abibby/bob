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
		{
			Name: "add foreign key",
			Builder: schema.Table("foo", func(table *schema.Blueprint) {
				table.ForeignKey("id", "bar", "foo_id")
			}),
			ExpectedSQL:      "CONSTRAINT \"foo\"(FOREIGN KEY (\"id\") REFERENCES \"bar\"(\"foo_id\"));",
			ExpectedBindings: []any{},
		},
		// {
		// 	Name: "drop foreign key",
		// 	Builder: schema.Table("foo", func(table *schema.Blueprint) {
		// 		table.ForeignKey("id", "bar", "foo_id")
		// 	}),
		// 	ExpectedSQL:      "",
		// 	ExpectedBindings: []any{},
		// },
		{
			Name: "add index",
			Builder: schema.Table("foo", func(table *schema.Blueprint) {
				table.Index("index-name").AddColumn("foo").AddColumn("bar")
			}),
			ExpectedSQL:      "CREATE INDEX IF NOT EXIST \"index-name\" ON \"foo\" (\"foo\", \"bar\");",
			ExpectedBindings: []any{},
		},
		// {
		// 	Name: "drop index",
		// 	Builder: schema.Table("foo", func(table *schema.Blueprint) {
		// 		table.Index("index-name").AddColumn("foo").AddColumn("bar")
		// 	}),
		// 	ExpectedSQL:      "",
		// 	ExpectedBindings: []any{},
		// },
	})
}
