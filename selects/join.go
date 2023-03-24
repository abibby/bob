package selects

type Join struct{}

func (b *Builder) Join() *Builder {
	return b
}
