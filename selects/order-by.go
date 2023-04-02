package selects

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type orderBys []builder.ToSQLer

func (o orderBys) Clone() orderBys {
	return cloneSlice(o)
}
func (o orderBys) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	if len(o) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("ORDER BY")
	r.Add(builder.Join(o, ", ").ToSQL(ctx, d))
	return r.ToSQL(ctx, d)
}

func (o orderBys) OrderBy(column string) orderBys {
	return append(o, builder.Identifier(column))
}
func (o orderBys) OrderByDesc(column string) orderBys {
	return append(o, builder.Join([]builder.ToSQLer{builder.Identifier(column), builder.Raw("DESC")}, " "))
}
