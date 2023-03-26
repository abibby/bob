package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type groupBys []builder.ToSQLer

func (g groupBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("GROUP BY")
	r.Add(builder.Join(g, ", ").ToSQL(d))
	return r.ToSQL(d)
}
func (b *Builder[T]) GroupBy(columns ...string) *Builder[T] {
	b.groupBys = builder.IdentifierList(columns)
	return b
}

func (b *Builder[T]) AddGroupBy(columns ...string) *Builder[T] {
	b.groupBys = append(b.groupBys, builder.IdentifierList(columns)...)
	return b
}
