package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type where struct {
	Column   string
	Operator string
	Value    any
	Or       bool
}

type whereList struct {
	list []*where
}

func newWhereList() *whereList {
	return &whereList{
		list: []*where{},
	}
}

func (w *whereList) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(w.list) == 0 {
		return "", nil, nil
	}

	r := builder.Result()
	for i, w := range w.list {
		if i != 0 {
			if w.Or == true {
				r.AddString("OR")
			} else {
				r.AddString("AND")
			}
		}
		if w.Column != "" {
			r.AddString(d.Identifier(w.Column))
		}
		if w.Operator != "" {
			r.AddString(w.Operator)
		}
		if sb, ok := w.Value.(*Builder); ok {
			r.Add(builder.NewGroup(sb).ToSQL(d))
		} else if sb, ok := w.Value.(*whereList); ok {
			r.Add(builder.NewGroup(sb).ToSQL(d))
		} else if sb, ok := w.Value.(builder.ToSQLer); ok {
			r.Add(sb.ToSQL(d))
		} else {
			r.Add(builder.NewLiteral(w.Value).ToSQL(d))
		}
	}

	return r.ToSQL(d)
}

func (w *whereList) Where(column, operator string, value any) *whereList {
	return w.where(column, operator, value, false)
}

func (w *whereList) OrWhere(column, operator string, value any) *whereList {
	return w.where(column, operator, value, true)
}

func (w *whereList) WhereIn(column string, values []any) *whereList {
	return w.whereIn(column, values, false)
}

func (w *whereList) OrWhereIn(column string, values []any) *whereList {
	return w.whereIn(column, values, true)
}

func (w *whereList) whereIn(column string, values []any, or bool) *whereList {
	return w.where(column, "in", builder.NewGroup(builder.Join(builder.LiteralList(values), ", ")), or)
}

func (w *whereList) where(column, operator string, value any, or bool) *whereList {
	w.list = append(w.list, &where{
		Column:   column,
		Operator: operator,
		Value:    value,
		Or:       or,
	})
	return w
}

func (w *whereList) And(cb func(w *whereList)) *whereList {
	subBuilder := newWhereList()
	cb(subBuilder)
	w.list = append(w.list, &where{Value: subBuilder})
	return w
}

func (w *whereList) Or(cb func(w *whereList)) *whereList {
	subBuilder := newWhereList()
	cb(subBuilder)
	w.list = append(w.list, &where{Value: subBuilder, Or: true})
	return w
}
