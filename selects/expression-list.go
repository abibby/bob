package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ExpressionList []builder.ToSQLer

func (e ExpressionList) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	for i, expr := range e {
		q, bindings, err := expr.ToSQL(d)
		if i < len(e)-1 {
			q += ","
		}
		r.Add(q, bindings, err)
	}
	return r.ToSQL(d)
}
