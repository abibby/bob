package craetetable

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestBuilder(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "create table",
			Builder:          CreateTable("foo", func(t *Table) {}),
			ExpectedSQL:      "CREATE TABLE \"foo\" ( )",
			ExpectedBindings: []any{},
		},
		{
			Name: "1 column",
			Builder: CreateTable("foo", func(t *Table) {
				t.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" ( \"bar\" text )",
			ExpectedBindings: []any{},
		},
		{
			Name: "2 columns",
			Builder: CreateTable("foo", func(t *Table) {
				t.Int("id")
				t.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" ( \"id\" int, \"bar\" text )",
			ExpectedBindings: []any{},
		},
	})
}
