package selects

import (
	"github.com/abibby/bob/dialects"
	"github.com/davecgh/go-spew/spew"
)

func (b *Builder) Dump(d dialects.Dialect) *Builder {
	spew.Dump(b.ToSQL(d))
	return b
}
