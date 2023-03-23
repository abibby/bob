package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type Wheres struct {
	*whereList
}

func NewWheres() *Wheres {
	return &Wheres{
		whereList: newWhereList(),
	}
}

func (w *Wheres) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := builder.Result()
	r.AddString("WHERE")
	r.Add(w.whereList.ToSQL(d))
	return r.ToSQL(d)
}
func (b *Builder) Where(column, operator string, value any) *Builder {
	b.wheres.whereList = b.wheres.Where(column, operator, value)
	return b
}
func (b *Builder) OrWhere(column, operator string, value any) *Builder {
	b.wheres.whereList = b.wheres.OrWhere(column, operator, value)
	return b
}
func (b *Builder) WhereIn(column string, values []any) *Builder {
	b.wheres.whereList = b.wheres.WhereIn(column, values)
	return b
}
func (b *Builder) OrWhereIn(column string, values []any) *Builder {
	b.wheres.whereList = b.wheres.OrWhereIn(column, values)
	return b
}
func (b *Builder) And(cb func(b *whereList)) *Builder {
	b.wheres.whereList = b.wheres.And(cb)
	return b
}

func (b *Builder) Or(cb func(b *whereList)) *Builder {
	b.wheres.whereList = b.wheres.Or(cb)
	return b
}
