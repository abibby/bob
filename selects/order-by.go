package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type OrderBys ExpressionList

func (g OrderBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("ORDER BY")
	r.Add(ExpressionList(g).ToSQL(d))
	return r.ToSQL(d)
}
func (b *Builder) OrderBy(columns ...string) *Builder {
	b.orderBys = builder.IdentifierList(columns)
	return b
}

func (b *Builder) AddOrderBy(columns ...string) *Builder {
	b.orderBys = append(b.orderBys, builder.IdentifierList(columns)...)
	return b
}
