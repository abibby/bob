package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type groupBys []builder.ToSQLer

func (g groupBys) Clone() groupBys {
	return cloneSlice(g)
}
func (g groupBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("GROUP BY")
	r.Add(builder.Join(g, ", "))
	return r.ToSQL(d)
}

func (b groupBys) GroupBy(columns ...string) groupBys {
	return builder.IdentifierList(columns)
}

func (b groupBys) AddGroupBy(columns ...string) groupBys {
	return append(b, builder.IdentifierList(columns)...)
}
