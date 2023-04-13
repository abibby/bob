package selects_test

import (
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
)

func TestJoin(t *testing.T) {
	test.QueryTest(t, []test.Case{
		{
			Name:             "join",
			Builder:          NewTestBuilder().Join("bars", "bars.foo_id", "=", "foos.id"),
			ExpectedSQL:      "SELECT * FROM \"foos\" JOIN \"bars\" ON \"bars\".\"foo_id\" = \"foos\".\"id\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "left join",
			Builder:          NewTestBuilder().LeftJoin("bars", "bars.foo_id", "=", "foos.id"),
			ExpectedSQL:      "SELECT * FROM \"foos\" LEFT JOIN \"bars\" ON \"bars\".\"foo_id\" = \"foos\".\"id\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "right join",
			Builder:          NewTestBuilder().RightJoin("bars", "bars.foo_id", "=", "foos.id"),
			ExpectedSQL:      "SELECT * FROM \"foos\" RIGHT JOIN \"bars\" ON \"bars\".\"foo_id\" = \"foos\".\"id\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "inner join",
			Builder:          NewTestBuilder().InnerJoin("bars", "bars.foo_id", "=", "foos.id"),
			ExpectedSQL:      "SELECT * FROM \"foos\" INNER JOIN \"bars\" ON \"bars\".\"foo_id\" = \"foos\".\"id\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "cross join",
			Builder:          NewTestBuilder().CrossJoin("bars", "bars.foo_id", "=", "foos.id"),
			ExpectedSQL:      "SELECT * FROM \"foos\" CROSS JOIN \"bars\" ON \"bars\".\"foo_id\" = \"foos\".\"id\"",
			ExpectedBindings: []any{},
		},
		{
			Name: "join on",
			Builder: NewTestBuilder().JoinOn("bars", func(q *selects.Conditions) {
				q.Where("a", ">", 4).WhereColumn("b", "=", "c")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" JOIN \"bars\" ON \"a\" > ? AND \"b\" = \"c\"",
			ExpectedBindings: []any{4},
		},
		{
			Name: "left join on",
			Builder: NewTestBuilder().LeftJoinOn("bars", func(q *selects.Conditions) {
				q.Where("a", ">", 4).WhereColumn("b", "=", "c")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" LEFT JOIN \"bars\" ON \"a\" > ? AND \"b\" = \"c\"",
			ExpectedBindings: []any{4},
		},
		{
			Name: "right join on",
			Builder: NewTestBuilder().RightJoinOn("bars", func(q *selects.Conditions) {
				q.Where("a", ">", 4).WhereColumn("b", "=", "c")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" RIGHT JOIN \"bars\" ON \"a\" > ? AND \"b\" = \"c\"",
			ExpectedBindings: []any{4},
		},
		{
			Name: "inner join on",
			Builder: NewTestBuilder().InnerJoinOn("bars", func(q *selects.Conditions) {
				q.Where("a", ">", 4).WhereColumn("b", "=", "c")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" INNER JOIN \"bars\" ON \"a\" > ? AND \"b\" = \"c\"",
			ExpectedBindings: []any{4},
		},
		{
			Name: "cross join on",
			Builder: NewTestBuilder().CrossJoinOn("bars", func(q *selects.Conditions) {
				q.Where("a", ">", 4).WhereColumn("b", "=", "c")
			}),
			ExpectedSQL:      "SELECT * FROM \"foos\" CROSS JOIN \"bars\" ON \"a\" > ? AND \"b\" = \"c\"",
			ExpectedBindings: []any{4},
		},
	})
}
