package bob

import (
	"testing"
)

func TestGroupBy(t *testing.T) {
	type Foo struct {
		A string `db:"a"`
	}
	foo := &Foo{}
	QueryTest(t, []TestCase[Foo]{
		{
			"one group",
			From(foo).GroupBy("a"),
			"SELECT * FROM `Foo` GROUP BY `a`",
			[]any{},
		},
		{
			"two groups",
			From(foo).GroupBy("a", "b"),
			"SELECT * FROM `Foo` GROUP BY `a`, `b`",
			[]any{},
		},
		{
			"different table",
			From(foo).GroupBy("a.b"),
			"SELECT * FROM `Foo` GROUP BY `a`.`b`",
			[]any{},
		},
	})
}
