package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestSelect(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"one select",
			NewTestBuilder().Select("a"),
			"SELECT `a` FROM `foo`",
			[]any{},
		},
		{
			"two select",
			NewTestBuilder().Select("a", "b"),
			"SELECT `a`, `b` FROM `foo`",
			[]any{},
		},
		{
			"different table",
			NewTestBuilder().Select("a.b"),
			"SELECT `a`.`b` FROM `foo`",
			[]any{},
		},
		{
			"distinct",
			NewTestBuilder().Select("a").Distinct(),
			"SELECT DISTINCT `a` FROM `foo`",
			[]any{},
		},
		{
			"subquery",
			NewTestBuilder().SelectSubquery(NewTestBuilder().Select("a")),
			"SELECT (SELECT `a` FROM `foo`) FROM `foo`",
			[]any{},
		},
	})
}
