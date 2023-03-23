package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type orderBys []builder.ToSQLer

func (g orderBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("ORDER BY")
	r.Add(builder.Join(g, ", ").ToSQL(d))
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
