package selects

import (
	"github.com/abibby/bob/dialects"
	"github.com/davecgh/go-spew/spew"
)

func (b *Builder[T]) Dump() *Builder[T] {
	spew.Dump(b.ToSQL(dialects.DefaultDialect))
	return b
}
