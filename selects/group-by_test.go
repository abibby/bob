package selects_test

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestGroupBy(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "one group",
			Builder:          NewTestBuilder().GroupBy("a"),
			ExpectedSQL:      "SELECT * FROM \"foos\" GROUP BY \"a\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "two groups",
			Builder:          NewTestBuilder().GroupBy("a", "b"),
			ExpectedSQL:      "SELECT * FROM \"foos\" GROUP BY \"a\", \"b\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "different table",
			Builder:          NewTestBuilder().GroupBy("a.b"),
			ExpectedSQL:      "SELECT * FROM \"foos\" GROUP BY \"a\".\"b\"",
			ExpectedBindings: []any{},
		},
	})
}
