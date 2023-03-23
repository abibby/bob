package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type GroupBys ExpressionList

func (g GroupBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("GROUP BY")
	r.Add(ExpressionList(g).ToSQL(d))
	return r.ToSQL(d)
}
func (b *Builder) GroupBy(columns ...string) *Builder {
	b.groupBys = builder.IdentifierList(columns)
	return b
}

func (b *Builder) AddGroupBy(columns ...string) *Builder {
	b.groupBys = append(b.groupBys, builder.IdentifierList(columns)...)
	return b
}
