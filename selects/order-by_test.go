package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestOrderBy(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "one group",
			Builder:          NewTestBuilder().OrderBy("a"),
			ExpectedSQL:      "SELECT * FROM \"foo\" ORDER BY \"a\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "two groups",
			Builder:          NewTestBuilder().OrderBy("a", "b"),
			ExpectedSQL:      "SELECT * FROM \"foo\" ORDER BY \"a\", \"b\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "different table",
			Builder:          NewTestBuilder().OrderBy("a.b"),
			ExpectedSQL:      "SELECT * FROM \"foo\" ORDER BY \"a\".\"b\"",
			ExpectedBindings: []any{},
		},
	})
}
