package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestGroupBy(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"one group",
			NewTestBuilder().GroupBy("a"),
			"SELECT * FROM `foo` GROUP BY `a`",
			[]any{},
		},
		{
			"two groups",
			NewTestBuilder().GroupBy("a", "b"),
			"SELECT * FROM `foo` GROUP BY `a`, `b`",
			[]any{},
		},
		{
			"different table",
			NewTestBuilder().GroupBy("a.b"),
			"SELECT * FROM `foo` GROUP BY `a`.`b`",
			[]any{},
		},
	})
}
