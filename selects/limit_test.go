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
			ExpectedSQL:      "SELECT * FROM \"foo\" LIMIT 1",
			ExpectedBindings: []any{},
		},
		{
			Name:             "offset",
			Builder:          NewTestBuilder().Offset(1),
			ExpectedSQL:      "SELECT * FROM \"foo\" LIMIT 0 OFFSET 1",
			ExpectedBindings: []any{},
		},
		{
			Name:             "limit and offset",
			Builder:          NewTestBuilder().Limit(1).Offset(2),
			ExpectedSQL:      "SELECT * FROM \"foo\" LIMIT 1 OFFSET 2",
			ExpectedBindings: []any{},
		},
	})
}
