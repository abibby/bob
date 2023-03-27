package selects

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type where struct {
	Column   string
	Operator string
	Value    any
	Or       bool
}

type WhereList struct {
	list []*where
}

func NewWhereList() *WhereList {
	return &WhereList{
		list: []*where{},
	}
}

func (w *WhereList) Clone() *WhereList {
	return &WhereList{
		list: cloneSlice(w.list),
	}
}
func (w *WhereList) ToSQL(d dialects.Dialect) (string, []any, error) {
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

		if w.Value == nil {
			switch w.Operator {
			case "=":
				r.AddString("IS NULL")
			case "!=":
				r.AddString("IS NOT NULL")
			default:
				return "", nil, fmt.Errorf("wheres checking nil only support = and !=")
			}
		} else {
			if w.Operator != "" {
				r.AddString(w.Operator)
			}
			if sb, ok := w.Value.(iBuilder); ok {
				r.Add(builder.NewGroup(sb).ToSQL(d))
			} else if sb, ok := w.Value.(*WhereList); ok {
				r.Add(builder.NewGroup(sb).ToSQL(d))
			} else if sb, ok := w.Value.(builder.ToSQLer); ok {
				r.Add(sb.ToSQL(d))
			} else {
				r.Add(builder.NewLiteral(w.Value).ToSQL(d))
			}
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

func (w *WhereList) WhereColumn(column, operator string, valueColumn string) *WhereList {
	return w.where(column, operator, builder.Identifier(valueColumn), false)
}

func (w *WhereList) OrWhereColumn(column, operator string, valueColumn string) *WhereList {
	return w.where(column, operator, builder.Identifier(valueColumn), true)
}

func (w *WhereList) WhereIn(column string, values []any) *WhereList {
	return w.whereIn(column, values, false)
}

func (w *WhereList) OrWhereIn(column string, values []any) *WhereList {
	return w.whereIn(column, values, true)
}

func (w *WhereList) whereIn(column string, values []any, or bool) *WhereList {
	return w.where(column, "in", builder.NewGroup(builder.Join(builder.LiteralList(values), ", ")), or)
}

func (w *WhereList) where(column, operator string, value any, or bool) *WhereList {
	w.list = append(w.list, &where{
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
	w.list = append(w.list, &where{Value: subBuilder})
	return w
}

func (w *WhereList) Or(cb func(w *WhereList)) *WhereList {
	subBuilder := NewWhereList()
	cb(subBuilder)
	w.list = append(w.list, &where{Value: subBuilder, Or: true})
	return w
}
