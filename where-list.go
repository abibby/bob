package bob

import (
	"github.com/abibby/bob/dialects"
)

type Where struct {
	Column   string
	Operator string
	Value    any
	Or       bool
}

type WhereList struct {
	list []*Where
}

func NewWhereList() *WhereList {
	return &WhereList{
		list: []*Where{},
	}
}

func (w *WhereList) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := &sqlResult{}
	for i, w := range w.list {
		if i != 0 {
			if w.Or == true {
				r.addString("OR")
			} else {
				r.addString("AND")
			}
		}
		if w.Column != "" {
			r.addString(d.Identifier(w.Column))
		}
		if w.Operator != "" {
			r.addString(w.Operator)
		}
		if sb, ok := w.Value.(*Builder); ok {
			r.add(NewGroup(sb).ToSQL(d))
		} else if sb, ok := w.Value.(*WhereList); ok {
			r.add(NewGroup(sb).ToSQL(d))
		} else if sb, ok := w.Value.(ToSQLer); ok {
			r.add(sb.ToSQL(d))
		} else {
			r.add(NewLiteral(w.Value).ToSQL(d))
		}
	}

	return r.ToSQL(d)
}

func (w *WhereList) Where(column, operator string, value any) *WhereList {
	return w.where(column, operator, value, false)
}

func (w *WhereList) OrWhere(column, operator string, value any) *WhereList {
	return w.where(column, operator, value, true)
}

func (w *WhereList) where(column, operator string, value any, or bool) *WhereList {
	w.list = append(w.list, &Where{
		Column:   column,
		Operator: operator,
		Value:    value,
		Or:       or,
	})
	return w
}

func (w *WhereList) And(cb func(w *WhereList)) *WhereList {
	subBuilder := NewWhereList()
	cb(subBuilder)
	w.list = append(w.list, &Where{Value: subBuilder})
	return w
}

func (w *WhereList) Or(cb func(w *WhereList)) *WhereList {
	subBuilder := NewWhereList()
	cb(subBuilder)
	w.list = append(w.list, &Where{Value: subBuilder, Or: true})
	return w
}
