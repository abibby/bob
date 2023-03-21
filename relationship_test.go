package bob

import (
	"testing"

	"github.com/abibby/bob/test"
)

type Foo struct {
	ID int `db:"id"`
}

func (f *Foo) Bar() *HasOne[*Foo, *Bar] {
	return NewHasOne(f, &Bar{}, "foo_id", "id")
}

type Bar struct {
	FooID int `db:"foo_id"`
}

func (f *Bar) Foo() *BelongsTo[*Foo] {
	return NewBelongsTo(&Foo{})
}

func TestHasOne(t *testing.T) {
	foo := &Foo{
		ID: 10,
	}
	test.QueryTest(t, []test.Case{
		{
			"has one",
			foo.Bar(),
			"SELECT * FROM `Bar` WHERE `foo_id` = ? LIMIT 1",
			[]any{foo.ID},
		},
	})
}
