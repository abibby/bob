package selects

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type where struct {
	Column   builder.ToSQLer
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
		if w.Column != nil {
			r.Add(w.Column.ToSQL(d))

			if w.Operator == "" {
				return "", nil, fmt.Errorf("the operator must be set when the column is set")
			}
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
			if sb, ok := w.Value.(QueryBuilder); ok {
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

func (w *WhereList) WhereExists(query QueryBuilder) *WhereList {
	return w.whereExists(query, false)
}
func (w *WhereList) OrWhereExists(query QueryBuilder) *WhereList {
	return w.whereExists(query, true)
}
func (w *WhereList) whereExists(query QueryBuilder, or bool) *WhereList {
	return w.addWhere(&where{
		Value: builder.ToSQLFunc(func(d dialects.Dialect) (string, []any, error) {
			return builder.Result().AddString("EXISTS").Add(query.ToSQL(d)).ToSQL(d)
		}),
		Or: or,
	})
}
func (w *WhereList) WhereSubquery(subquery QueryBuilder, operator string, value any) *WhereList {
	return w.whereSubquery(subquery, operator, value, false)
}
func (w *WhereList) OrWhereSubquery(subquery QueryBuilder, operator string, value any) *WhereList {
	return w.whereSubquery(subquery, operator, value, true)
}
func (w *WhereList) whereSubquery(subquery QueryBuilder, operator string, value any, or bool) *WhereList {
	return w.addWhere(&where{
		Column:   builder.NewGroup(subquery),
		Operator: operator,
		Value:    value,
		Or:       or,
	})
}

func (w *WhereList) where(column, operator string, value any, or bool) *WhereList {
	return w.addWhere(&where{
		Column:   builder.Identifier(column),
		Operator: operator,
		Value:    value,
		Or:       or,
	})
}
func (w *WhereList) addWhere(wh *where) *WhereList {
	w.list = append(w.list, wh)
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
