package bob

import (
	"testing"
)

func TestLimit(t *testing.T) {
	type Foo struct {
		A string `db:"a"`
	}
	foo := &Foo{}
	QueryTest(t, []TestCase[Foo]{
		{
			"limit",
			From(foo).Limit(1),
			"SELECT * FROM `Foo` LIMIT 1",
			[]any{},
		},
		{
			"offset",
			From(foo).Offset(1),
			"SELECT * FROM `Foo` LIMIT 0 OFFSET 1",
			[]any{},
		},
		{
			"limit and offset",
			From(foo).Limit(1).Offset(2),
			"SELECT * FROM `Foo` LIMIT 1 OFFSET 2",
			[]any{},
		},
	})
}
