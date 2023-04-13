package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type join struct {
	direction  string
	table      builder.ToSQLer
	conditions *Conditions
}

type joins []*join

func (j joins) Clone() joins {
	return cloneSlice(j)
}
func (j joins) Join(table, localColumn, operator, foreignColumn string) joins {
	return j.join("", table, localColumn, operator, foreignColumn)
}
func (j joins) LeftJoin(table, localColumn, operator, foreignColumn string) joins {
	return j.join("LEFT", table, localColumn, operator, foreignColumn)
}
func (j joins) RightJoin(table, localColumn, operator, foreignColumn string) joins {
	return j.join("RIGHT", table, localColumn, operator, foreignColumn)
}
func (j joins) InnerJoin(table, localColumn, operator, foreignColumn string) joins {
	return j.join("INNER", table, localColumn, operator, foreignColumn)
}
func (j joins) CrossJoin(table, localColumn, operator, foreignColumn string) joins {
	return j.join("CROSS", table, localColumn, operator, foreignColumn)
}
func (j joins) join(direction, table, localColumn, operator, foreignColumn string) joins {
	return j.joinOn(direction, table, func(q *Conditions) {
		q.WhereColumn(localColumn, operator, foreignColumn)
	})
}
func (j joins) JoinOn(table string, cb func(q *Conditions)) joins {
	return j.joinOn("", table, cb)
}
func (j joins) LeftJoinOn(table string, cb func(q *Conditions)) joins {
	return j.joinOn("LEFT", table, cb)
}
func (j joins) RightJoinOn(table string, cb func(q *Conditions)) joins {
	return j.joinOn("RIGHT", table, cb)
}
func (j joins) InnerJoinOn(table string, cb func(q *Conditions)) joins {
	return j.joinOn("INNER", table, cb)
}
func (j joins) CrossJoinOn(table string, cb func(q *Conditions)) joins {
	return j.joinOn("CROSS", table, cb)
}
func (j joins) joinOn(direction string, table string, cb func(q *Conditions)) joins {
	c := newConditions().withPrefix("ON")
	cb(c)
	return append(j, &join{
		direction:  direction,
		table:      builder.Identifier(table),
		conditions: c,
	})
}
func (j *join) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result().
		AddString(j.direction).
		AddString("JOIN").
		Add(j.table.ToSQL(d)).
		Add(j.conditions.ToSQL(d))

	return r.ToSQL(d)
}
func (j joins) ToSQL(d dialects.Dialect) (string, []any, error) {
	return builder.Join(j, " ").ToSQL(d)
}
