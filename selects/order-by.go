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
func (b *Builder[T]) OrderBy(columns ...string) *Builder[T] {
	b.orderBys = builder.IdentifierList(columns)
	return b
}

func (b *Builder[T]) AddOrderBy(columns ...string) *Builder[T] {
	b.orderBys = append(b.orderBys, builder.IdentifierList(columns)...)
	return b
}
