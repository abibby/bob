package craetetable_test

import (
	"testing"

	"github.com/abibby/bob/craetetable"
	"github.com/abibby/bob/test"
)

func TestBuilder(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "create table",
			Builder:          craetetable.CreateTable("foo", func(t *craetetable.Table) {}),
			ExpectedSQL:      "CREATE TABLE \"foo\" ()",
			ExpectedBindings: []any{},
		},
		{
			Name: "1 column",
			Builder: craetetable.CreateTable("foo", func(t *craetetable.Table) {
				t.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"bar\" TEXT NOT NULL)",
			ExpectedBindings: []any{},
		},
		{
			Name: "2 columns",
			Builder: craetetable.CreateTable("foo", func(t *craetetable.Table) {
				t.Int("id")
				t.String("bar")
			}),
			ExpectedSQL:      "CREATE TABLE \"foo\" (\"id\" INTEGER NOT NULL, \"bar\" TEXT NOT NULL)",
			ExpectedBindings: []any{},
		},
	})
}
