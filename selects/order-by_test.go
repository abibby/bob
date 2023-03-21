package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestOrderBy(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"one group",
			NewTestBuilder().OrderBy("a"),
			"SELECT * FROM `foo` ORDER BY `a`",
			[]any{},
		},
		{
			"two groups",
			NewTestBuilder().OrderBy("a", "b"),
			"SELECT * FROM `foo` ORDER BY `a`, `b`",
			[]any{},
		},
		{
			"different table",
			NewTestBuilder().OrderBy("a.b"),
			"SELECT * FROM `foo` ORDER BY `a`.`b`",
			[]any{},
		},
	})
}
