package selects

func (b *Builder[T]) Clone() *Builder[T] {
	return &Builder[T]{
		selects:  b.selects.Clone(),
		from:     b.from.Clone(),
		wheres:   b.wheres.Clone(),
		groupBys: b.groupBys.Clone(),
		havings:  b.havings.Clone(),
		limit:    b.limit.Clone(),
		orderBys: b.orderBys.Clone(),
	}
}

func cloneSlice[T any](arr []T) []T {
	l := make([]T, len(arr))
	for i, v := range arr {
		l[i] = v
	}
	return l
}
