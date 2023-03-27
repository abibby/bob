package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type orderBys []builder.ToSQLer

func (o orderBys) Clone() orderBys {
	return cloneSlice(o)
}
func (o orderBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(o) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("ORDER BY")
	r.Add(builder.Join(o, ", ").ToSQL(d))
	return r.ToSQL(d)
}

func (b *Builder[T]) OrderBy(column string) *Builder[T] {
	b.orderBys = append(b.orderBys, builder.Identifier(column))
	return b
}
func (b *Builder[T]) OrderByDesc(column string) *Builder[T] {
	b.orderBys = append(b.orderBys, builder.Join([]builder.ToSQLer{builder.Identifier(column), builder.Raw("DESC")}, " "))
	return b
}
