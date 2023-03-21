package selects

func NewTestBuilder() *Builder {
	return New().Select("*").From("foo")
}
