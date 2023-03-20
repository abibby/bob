package bob

import (
	"testing"
)

func TestHaving(t *testing.T) {
	type M struct {
		A string `db:"a"`
	}
	m := &M{}

	QueryTest(t, []TestCase[M]{
		{
			"one where",
			From(m).Having("a", "=", "b"),
			"SELECT * FROM `M` HAVING `a` = ?",
			[]any{"b"},
		},
		{
			"2 wheres",
			From(m).Having("a", "=", "b").Having("c", "=", "d"),
			"SELECT * FROM `M` HAVING `a` = ? AND `c` = ?",
			[]any{"b", "d"},
		},
		{
			"specified table",
			From(m).Having("foo.a", "=", "b"),
			"SELECT * FROM `M` HAVING `foo`.`a` = ?",
			[]any{"b"},
		},
		{
			"or where",
			From(m).Having("a", "=", "b").OrHaving("c", "=", "d"),
			"SELECT * FROM `M` HAVING `a` = ? OR `c` = ?",
			[]any{"b", "d"},
		},
		{
			"and group",
			From(m).HavingAnd(func(b *WhereList) {
				b.Where("a", "=", "a").OrWhere("b", "=", "b")
			}).HavingAnd(func(b *WhereList) {
				b.Where("c", "=", "c").OrWhere("d", "=", "d")
			}),
			"SELECT * FROM `M` HAVING (`a` = ? OR `b` = ?) AND (`c` = ? OR `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"or group",
			From(m).HavingOr(func(b *WhereList) {
				b.Where("a", "=", "a").Where("b", "=", "b")
			}).HavingOr(func(b *WhereList) {
				b.Where("c", "=", "c").Where("d", "=", "d")
			}),
			"SELECT * FROM `M` HAVING (`a` = ? AND `b` = ?) OR (`c` = ? AND `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"subquery",
			From(m).Having("a", "=", From(m).Select("a").Having("id", "=", 1)),
			"SELECT * FROM `M` HAVING `a` = (SELECT `a` FROM `M` HAVING `id` = ?)",
			[]any{1},
		},
	})
}
