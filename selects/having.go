package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type havings struct {
	*WhereList
}

func newHavings() *havings {
	return &havings{
		WhereList: newWhereList(),
	}
}

func (w *havings) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := builder.Result()
	r.AddString("HAVING")
	r.Add(w.WhereList.ToSQL(d))
	return r.ToSQL(d)
}

func (b *Builder) Having(column, operator string, value any) *Builder {
	b.havings.WhereList = b.havings.Where(column, operator, value)
	return b
}
func (b *Builder) OrHaving(column, operator string, value any) *Builder {
	b.havings.WhereList = b.havings.OrWhere(column, operator, value)
	return b
}
func (b *Builder) HavingAnd(cb func(b *WhereList)) *Builder {
	b.havings.WhereList = b.havings.And(cb)
	return b
}

func (b *Builder) HavingOr(cb func(b *WhereList)) *Builder {
	b.havings.WhereList = b.havings.Or(cb)
	return b
}
