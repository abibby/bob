package selects

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

func (b *Builder[T]) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	return b.subBuilder.ToSQL(ctx, d)
}
func (b *SubBuilder) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	b = b.Clone()
	for _, scope := range b.scopes.allScopes() {
		b = scope.Apply(ctx, b)
	}
	return builder.Result().
		Add(b.selects.ToSQL(ctx, d)).
		Add(b.from.ToSQL(ctx, d)).
		Add(b.wheres.ToSQL(ctx, d)).
		Add(b.groupBys.ToSQL(ctx, d)).
		Add(b.havings.ToSQL(ctx, d)).
		Add(b.orderBys.ToSQL(ctx, d)).
		Add(b.limit.ToSQL(ctx, d)).
		ToSQL(ctx, d)
}
