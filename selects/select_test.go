package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestSelect(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "one select",
			Builder:          NewTestBuilder().Select("a"),
			ExpectedSQL:      "SELECT \"a\" FROM \"foo\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "two select",
			Builder:          NewTestBuilder().Select("a", "b"),
			ExpectedSQL:      "SELECT \"a\", \"b\" FROM \"foo\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "different table",
			Builder:          NewTestBuilder().Select("a.b"),
			ExpectedSQL:      "SELECT \"a\".\"b\" FROM \"foo\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "distinct",
			Builder:          NewTestBuilder().Select("a").Distinct(),
			ExpectedSQL:      "SELECT DISTINCT \"a\" FROM \"foo\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "subquery",
			Builder:          NewTestBuilder().SelectSubquery(NewTestBuilder().Select("a")),
			ExpectedSQL:      "SELECT (SELECT \"a\" FROM \"foo\") FROM \"foo\"",
			ExpectedBindings: []any{},
		},
	})
}
