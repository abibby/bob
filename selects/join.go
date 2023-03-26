package selects

type Join struct{}

func (b *Builder[T]) Join() *Builder[T] {
	return b
}
