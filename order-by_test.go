package bob

import (
	"testing"
)

func TestOrderBy(t *testing.T) {
	type Foo struct {
		A string `db:"a"`
	}
	foo := &Foo{}
	QueryTest(t, []TestCase[Foo]{
		{
			"one group",
			From(foo).OrderBy("a"),
			"SELECT * FROM `Foo` ORDER BY `a`",
			[]any{},
		},
		{
			"two groups",
			From(foo).OrderBy("a", "b"),
			"SELECT * FROM `Foo` ORDER BY `a`, `b`",
			[]any{},
		},
		{
			"different table",
			From(foo).OrderBy("a.b"),
			"SELECT * FROM `Foo` ORDER BY `a`.`b`",
			[]any{},
		},
	})
}
