package selects

import (
	"github.com/abibby/bob/dialects"
	"github.com/davecgh/go-spew/spew"
)

func (b *SubBuilder) Dump() *SubBuilder {
	spew.Dump(b.ToSQL(dialects.DefaultDialect))
	return b
}
