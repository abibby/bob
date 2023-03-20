package bob

import "github.com/abibby/bob/dialects"

type Havings struct {
	*WhereList
}

func NewHavings() *Havings {
	return &Havings{
		WhereList: NewWhereList(),
	}
}

func (w *Havings) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := &sqlResult{}
	r.addString("HAVING")
	r.add(w.WhereList.ToSQL(d))
	return r.ToSQL(d)
}

func (b *SelectBuilder) Having(column, operator string, value any) *SelectBuilder {
	b.havings.WhereList = b.havings.Where(column, operator, value)
	return b
}
func (b *SelectBuilder) OrHaving(column, operator string, value any) *SelectBuilder {
	b.havings.WhereList = b.havings.OrWhere(column, operator, value)
	return b
}
func (b *SelectBuilder) HavingAnd(cb func(b *WhereList)) *SelectBuilder {
	b.havings.WhereList = b.havings.And(cb)
	return b
}

func (b *SelectBuilder) HavingOr(cb func(b *WhereList)) *SelectBuilder {
	b.havings.WhereList = b.havings.Or(cb)
	return b
}
