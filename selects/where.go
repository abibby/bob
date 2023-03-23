package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type Wheres struct {
	*WhereList
}

func NewWheres() *Wheres {
	return &Wheres{
		WhereList: NewWhereList(),
	}
}

func (w *Wheres) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := builder.Result()
	r.AddString("WHERE")
	r.Add(w.WhereList.ToSQL(d))
	return r.ToSQL(d)
}
func (b *Builder) Where(column, operator string, value any) *Builder {
	b.wheres.WhereList = b.wheres.Where(column, operator, value)
	return b
}
func (b *Builder) OrWhere(column, operator string, value any) *Builder {
	b.wheres.WhereList = b.wheres.OrWhere(column, operator, value)
	return b
}
func (b *Builder) WhereIn(column string, values []any) *Builder {
	b.wheres.WhereList = b.wheres.WhereIn(column, values)
	return b
}
func (b *Builder) OrWhereIn(column string, values []any) *Builder {
	b.wheres.WhereList = b.wheres.OrWhereIn(column, values)
	return b
}
func (b *Builder) And(cb func(b *WhereList)) *Builder {
	b.wheres.WhereList = b.wheres.And(cb)
	return b
}

func (b *Builder) Or(cb func(b *WhereList)) *Builder {
	b.wheres.WhereList = b.wheres.Or(cb)
	return b
}
