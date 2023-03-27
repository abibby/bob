package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

func (b *Builder[T]) ToSQL(d dialects.Dialect) (string, []any, error) {
	return builder.Result().
		Add(b.selects.ToSQL(d)).
		Add(b.from.ToSQL(d)).
		Add(b.wheres.ToSQL(d)).
		Add(b.groupBys.ToSQL(d)).
		Add(b.havings.ToSQL(d)).
		Add(b.orderBys.ToSQL(d)).
		Add(b.limit.ToSQL(d)).
		ToSQL(d)
}
