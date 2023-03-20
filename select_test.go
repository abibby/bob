package bob

import (
	"testing"
)

func TestSelect(t *testing.T) {
	type Foo struct {
		A string `db:"a"`
	}
	foo := &Foo{}
	QueryTest(t, []TestCase[Foo]{
		{
			"one select",
			From(foo).Select("a"),
			"SELECT `a` FROM `Foo`",
			[]any{},
		},
		{
			"two select",
			From(foo).Select("a", "b"),
			"SELECT `a`, `b` FROM `Foo`",
			[]any{},
		},
		{
			"different table",
			From(foo).Select("a.b"),
			"SELECT `a`.`b` FROM `Foo`",
			[]any{},
		},
		{
			"distinct",
			From(foo).Select("a").Distinct(),
			"SELECT DISTINCT `a` FROM `Foo`",
			[]any{},
		},
		{
			"subquery",
			From(foo).SelectSubquery(From(foo).Select("a")),
			"SELECT (SELECT `a` FROM `Foo`) FROM `Foo`",
			[]any{},
		},
	})
}
