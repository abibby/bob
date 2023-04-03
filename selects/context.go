package selects

import "context"

func (b *SubBuilder) WithContext(ctx context.Context) *SubBuilder {
	b = b.Clone()
	b.ctx = ctx
	return b
}

func (b *SubBuilder) Context() context.Context {
	return b.ctx
}

func (b *Builder[T]) Context() context.Context {
	return b.subBuilder.Context()
}
