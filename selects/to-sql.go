package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

func (b *Builder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.Add(b.selects.ToSQL(d))
	r.Add(b.from.ToSQL(d))
	r.Add(b.wheres.ToSQL(d))
	r.Add(b.groupBys.ToSQL(d))
	r.Add(b.havings.ToSQL(d))
	r.Add(b.limit.ToSQL(d))
	r.Add(b.orderBys.ToSQL(d))

	return r.ToSQL(d)
}
