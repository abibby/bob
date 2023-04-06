package selects_test

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestLimit(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "limit",
			Builder:          NewTestBuilder().Limit(1),
			ExpectedSQL:      "SELECT * FROM \"foos\" LIMIT ?",
			ExpectedBindings: []any{1},
		},
		{
			Name:             "offset",
			Builder:          NewTestBuilder().Offset(1),
			ExpectedSQL:      "SELECT * FROM \"foos\" LIMIT ? OFFSET ?",
			ExpectedBindings: []any{0, 1},
		},
		{
			Name:             "limit and offset",
			Builder:          NewTestBuilder().Limit(1).Offset(2),
			ExpectedSQL:      "SELECT * FROM \"foos\" LIMIT ? OFFSET ?",
			ExpectedBindings: []any{1, 2},
		},
	})
}
