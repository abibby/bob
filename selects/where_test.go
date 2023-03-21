package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestWhere(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"one where",
			NewTestBuilder().Where("a", "=", "b"),
			"SELECT * FROM `foo` WHERE `a` = ?",
			[]any{"b"},
		},
		{
			"2 wheres",
			NewTestBuilder().Where("a", "=", "b").Where("c", "=", "d"),
			"SELECT * FROM `foo` WHERE `a` = ? AND `c` = ?",
			[]any{"b", "d"},
		},
		{
			"specified table",
			NewTestBuilder().Where("foo.a", "=", "b"),
			"SELECT * FROM `foo` WHERE `foo`.`a` = ?",
			[]any{"b"},
		},
		{
			"or where",
			NewTestBuilder().Where("a", "=", "b").OrWhere("c", "=", "d"),
			"SELECT * FROM `foo` WHERE `a` = ? OR `c` = ?",
			[]any{"b", "d"},
		},
		{
			"and group",
			NewTestBuilder().And(func(b *WhereList) {
				b.Where("a", "=", "a").OrWhere("b", "=", "b")
			}).And(func(b *WhereList) {
				b.Where("c", "=", "c").OrWhere("d", "=", "d")
			}),
			"SELECT * FROM `foo` WHERE (`a` = ? OR `b` = ?) AND (`c` = ? OR `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"or group",
			NewTestBuilder().Or(func(b *WhereList) {
				b.Where("a", "=", "a").Where("b", "=", "b")
			}).Or(func(b *WhereList) {
				b.Where("c", "=", "c").Where("d", "=", "d")
			}),
			"SELECT * FROM `foo` WHERE (`a` = ? AND `b` = ?) OR (`c` = ? AND `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"subquery",
			NewTestBuilder().Where("a", "=", NewTestBuilder().Select("a").Where("id", "=", 1)),
			"SELECT * FROM `foo` WHERE `a` = (SELECT `a` FROM `foo` WHERE `id` = ?)",
			[]any{1},
		},
	})
}
