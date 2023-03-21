package selects

import "github.com/abibby/bob/dialects"

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

	r := &sqlResult{}
	r.addString("WHERE")
	r.add(w.WhereList.ToSQL(d))
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
func (b *Builder) And(cb func(b *WhereList)) *Builder {
	b.wheres.WhereList = b.wheres.And(cb)
	return b
}

func (b *Builder) Or(cb func(b *WhereList)) *Builder {
	b.wheres.WhereList = b.wheres.Or(cb)
	return b
}
