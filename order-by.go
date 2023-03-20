package bob

import "github.com/abibby/bob/dialects"

type OrderBys ExpressionList

func (g OrderBys) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(g) == 0 {
		return "", nil, nil
	}
	r := &sqlResult{}
	r.addString("ORDER BY")
	r.add(ExpressionList(g).ToSQL(d))
	return r.ToSQL(d)
}
func (b *SelectBuilder) OrderBy(columns ...string) *SelectBuilder {
	b.orderBys = IdentifierList(columns)
	return b
}

func (b *SelectBuilder) AddOrderBy(columns ...string) *SelectBuilder {
	b.orderBys = append(b.orderBys, IdentifierList(columns)...)
	return b
}
