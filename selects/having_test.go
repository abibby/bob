package selects

import (
	"testing"

	"github.com/abibby/bob/test"
)

func TestHaving(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			"one where",
			NewTestBuilder().Having("a", "=", "b"),
			"SELECT * FROM `foo` HAVING `a` = ?",
			[]any{"b"},
		},
		{
			"2 wheres",
			NewTestBuilder().Having("a", "=", "b").Having("c", "=", "d"),
			"SELECT * FROM `foo` HAVING `a` = ? AND `c` = ?",
			[]any{"b", "d"},
		},
		{
			"specified table",
			NewTestBuilder().Having("foo.a", "=", "b"),
			"SELECT * FROM `foo` HAVING `foo`.`a` = ?",
			[]any{"b"},
		},
		{
			"or where",
			NewTestBuilder().Having("a", "=", "b").OrHaving("c", "=", "d"),
			"SELECT * FROM `foo` HAVING `a` = ? OR `c` = ?",
			[]any{"b", "d"},
		},
		{
			"and group",
			NewTestBuilder().HavingAnd(func(b *WhereList) {
				b.Where("a", "=", "a").OrWhere("b", "=", "b")
			}).HavingAnd(func(b *WhereList) {
				b.Where("c", "=", "c").OrWhere("d", "=", "d")
			}),
			"SELECT * FROM `foo` HAVING (`a` = ? OR `b` = ?) AND (`c` = ? OR `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"or group",
			NewTestBuilder().HavingOr(func(b *WhereList) {
				b.Where("a", "=", "a").Where("b", "=", "b")
			}).HavingOr(func(b *WhereList) {
				b.Where("c", "=", "c").Where("d", "=", "d")
			}),
			"SELECT * FROM `foo` HAVING (`a` = ? AND `b` = ?) OR (`c` = ? AND `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"subquery",
			NewTestBuilder().Having("a", "=", NewTestBuilder().Select("a").Having("id", "=", 1)),
			"SELECT * FROM `foo` HAVING `a` = (SELECT `a` FROM `foo` HAVING `id` = ?)",
			[]any{1},
		},
	})
}
