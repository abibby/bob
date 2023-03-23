package relationships

// make sure these match the models in test

type Foo struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Bar  *HasOne[*Bar]
	Bars *HasMany[*Bar]
}
type Bar struct {
	ID    int `db:"id"`
	FooID int `db:"foo_id"`
	Foo   *BelongsTo[*Foo]
}
