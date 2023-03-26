package selects_test

import (
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
)

func TestWhere(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "one where",
			Builder:          NewTestBuilder().Where("a", "=", "b"),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" = ?",
			ExpectedBindings: []any{"b"},
		},
		{
			Name:             "2 wheres",
			Builder:          NewTestBuilder().Where("a", "=", "b").Where("c", "=", "d"),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" = ? AND \"c\" = ?",
			ExpectedBindings: []any{"b", "d"},
		},
		{
			Name:             "specified table",
			Builder:          NewTestBuilder().Where("foo.a", "=", "b"),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"foo\".\"a\" = ?",
			ExpectedBindings: []any{"b"},
		},
		{
			Name:             "or where",
			Builder:          NewTestBuilder().Where("a", "=", "b").OrWhere("c", "=", "d"),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" = ? OR \"c\" = ?",
			ExpectedBindings: []any{"b", "d"},
		},
		{
			Name: "and group",
			Builder: NewTestBuilder().And(func(b *selects.WhereList) {
				b.Where("a", "=", "a").OrWhere("b", "=", "b")
			}).And(func(b *selects.WhereList) {
				b.Where("c", "=", "c").OrWhere("d", "=", "d")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE (\"a\" = ? OR \"b\" = ?) AND (\"c\" = ? OR \"d\" = ?)",
			ExpectedBindings: []any{"a", "b", "c", "d"},
		},
		{
			Name: "or group",
			Builder: NewTestBuilder().Or(func(b *selects.WhereList) {
				b.Where("a", "=", "a").Where("b", "=", "b")
			}).Or(func(b *selects.WhereList) {
				b.Where("c", "=", "c").Where("d", "=", "d")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE (\"a\" = ? AND \"b\" = ?) OR (\"c\" = ? AND \"d\" = ?)",
			ExpectedBindings: []any{"a", "b", "c", "d"},
		},
		{
			Name:             "subquery",
			Builder:          NewTestBuilder().Where("a", "=", NewTestBuilder().Select("a").Where("id", "=", 1)),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" = (SELECT \"a\" FROM \"foos\" WHERE \"id\" = ?)",
			ExpectedBindings: []any{1},
		},
		{
			Name:             "wherein",
			Builder:          NewTestBuilder().WhereIn("a", []any{1, 2, 3}),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" in (?, ?, ?)",
			ExpectedBindings: []any{1, 2, 3},
		},
	})
}
