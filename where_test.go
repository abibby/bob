package bob

import (
	"testing"
)

func TestWhere(t *testing.T) {
	type M struct {
		A string `db:"a"`
	}
	m := &M{}
	QueryTest(t, []TestCase[M]{
		{
			"one where",
			From(m).Where("a", "=", "b"),
			"SELECT * FROM `M` WHERE `a` = ?",
			[]any{"b"},
		},
		{
			"2 wheres",
			From(m).Where("a", "=", "b").Where("c", "=", "d"),
			"SELECT * FROM `M` WHERE `a` = ? AND `c` = ?",
			[]any{"b", "d"},
		},
		{
			"specified table",
			From(m).Where("foo.a", "=", "b"),
			"SELECT * FROM `M` WHERE `foo`.`a` = ?",
			[]any{"b"},
		},
		{
			"or where",
			From(m).Where("a", "=", "b").OrWhere("c", "=", "d"),
			"SELECT * FROM `M` WHERE `a` = ? OR `c` = ?",
			[]any{"b", "d"},
		},
		{
			"and group",
			From(m).And(func(b *WhereList) {
				b.Where("a", "=", "a").OrWhere("b", "=", "b")
			}).And(func(b *WhereList) {
				b.Where("c", "=", "c").OrWhere("d", "=", "d")
			}),
			"SELECT * FROM `M` WHERE (`a` = ? OR `b` = ?) AND (`c` = ? OR `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"or group",
			From(m).Or(func(b *WhereList) {
				b.Where("a", "=", "a").Where("b", "=", "b")
			}).Or(func(b *WhereList) {
				b.Where("c", "=", "c").Where("d", "=", "d")
			}),
			"SELECT * FROM `M` WHERE (`a` = ? AND `b` = ?) OR (`c` = ? AND `d` = ?)",
			[]any{"a", "b", "c", "d"},
		},
		{
			"subquery",
			From(m).Where("a", "=", From(m).Select("a").Where("id", "=", 1)),
			"SELECT * FROM `M` WHERE `a` = (SELECT `a` FROM `M` WHERE `id` = ?)",
			[]any{1},
		},
	})
}
