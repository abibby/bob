package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestLimit(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"limit",
			NewTestBuilder().Limit(1),
			"SELECT * FROM `foo` LIMIT 1",
			[]any{},
		},
		{
			"offset",
			NewTestBuilder().Offset(1),
			"SELECT * FROM `foo` LIMIT 0 OFFSET 1",
			[]any{},
		},
		{
			"limit and offset",
			NewTestBuilder().Limit(1).Offset(2),
			"SELECT * FROM `foo` LIMIT 1 OFFSET 2",
			[]any{},
		},
	})
}
