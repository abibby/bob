package schema_test

import (
	"testing"

	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/test"
)

func TestBuilder(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "create table",
			Builder:          schema.Create("foo", func(table *schema.Blueprint) {}),
			ExpectedSQL:      "CREATE TABLE \"foo\" ();",
			ExpectedBindings: []any{},
		},
		{
			Name: "1 column",
			Builder: schema.Create("foo", func(table *schema.Blueprint) {
				table.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"bar\" TEXT NOT NULL);",
			ExpectedBindings: []any{},
		},
		{
			Name: "2 columns",
			Builder: schema.Create("foo", func(table *schema.Blueprint) {
				table.Int("id")
				table.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"id\" INTEGER NOT NULL, \"bar\" TEXT NOT NULL);",
			ExpectedBindings: []any{},
		},
		{
			Name: "primary key",
			Builder: schema.Create("foo", func(table *schema.Blueprint) {
				table.Int("id").Primary()
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"id\" INTEGER NOT NULL, PRIMARY KEY (\"id\"));",
			ExpectedBindings: []any{},
		},
		{
			Name: "composite primary key",
			Builder: schema.Create("foo", func(table *schema.Blueprint) {
				table.Int("id1").Primary()
				table.Int("id2").Primary()
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"id1\" INTEGER NOT NULL, \"id2\" INTEGER NOT NULL, PRIMARY KEY (\"id1\", \"id2\"));",
			ExpectedBindings: []any{},
		},
	})
}
