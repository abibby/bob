package selects_test

import "github.com/abibby/bob/selects"

// make sure these match the models in test

type Foo struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Bar  *selects.HasOne[*Bar]
	Bars *selects.HasMany[*Bar]
}
type Bar struct {
	ID    int `db:"id"`
	FooID int `db:"foo_id"`
	Foo   *selects.BelongsTo[*Foo]
}
