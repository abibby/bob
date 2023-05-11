package schema_test

import (
	"testing"

	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/test"
)

func TestColumnBuilder(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "column",
			Builder:          schema.NewColumn("foo", dialects.DataTypeInt32),
			ExpectedSQL:      "\"foo\" INTEGER NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "Nullable",
			Builder:          schema.NewColumn("foo", dialects.DataTypeString).Nullable(),
			ExpectedSQL:      "\"foo\" TEXT",
			ExpectedBindings: []any{},
		},
		{
			Name:             "NotNullable",
			Builder:          schema.NewColumn("foo", dialects.DataTypeString).Nullable().NotNullable(),
			ExpectedSQL:      "\"foo\" TEXT NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "Primary",
			Builder:          schema.NewColumn("foo", dialects.DataTypeInt32).Primary(),
			ExpectedSQL:      "\"foo\" INTEGER NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "AutoIncrement",
			Builder:          schema.NewColumn("foo", dialects.DataTypeInt32).AutoIncrement(),
			ExpectedSQL:      "\"foo\" INTEGER AUTOINCREMENT NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "Default",
			Builder:          schema.NewColumn("foo", dialects.DataTypeString).Default("bar"),
			ExpectedSQL:      "\"foo\" TEXT NOT NULL DEFAULT ?",
			ExpectedBindings: []any{"bar"},
		},
		{
			Name:             "Type",
			Builder:          schema.NewColumn("foo", dialects.DataTypeString).Type(dialects.DataTypeInt32),
			ExpectedSQL:      "\"foo\" INTEGER NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "Unique",
			Builder:          schema.NewColumn("foo", dialects.DataTypeString).Unique(),
			ExpectedSQL:      "\"foo\" TEXT NOT NULL UNIQUE",
			ExpectedBindings: []any{},
		},
		{
			Name:             "DefaultCurrentTime",
			Builder:          schema.NewColumn("foo", dialects.DataTypeDateTime).DefaultCurrentTime(),
			ExpectedSQL:      "\"foo\" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP",
			ExpectedBindings: []any{},
		},
	})
}
