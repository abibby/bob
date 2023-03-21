package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ExpressionList []builder.ToSQLer

func (e ExpressionList) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := &sqlResult{}
	for i, expr := range e {
		q, args, err := expr.ToSQL(d)
		if i < len(e)-1 {
			q += ","
		}
		r.add(q, args, err)
	}
	return r.ToSQL(d)
}
