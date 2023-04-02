package selects

import (
	"context"

	"github.com/abibby/bob/dialects"
	"github.com/davecgh/go-spew/spew"
)

func (b *SubBuilder) Dump(ctx context.Context) *SubBuilder {
	spew.Dump(b.ToSQL(ctx, dialects.DefaultDialect))
	return b
}
