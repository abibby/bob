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
		WhereList: NewWhereList(),
	}
}

func (w *havings) Clone() *havings {
	return &havings{
		WhereList: w.WhereList.Clone(),
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

func (b *Builder[T]) Having(column, operator string, value any) *Builder[T] {
	b.havings.WhereList = b.havings.Where(column, operator, value)
	return b
}
func (b *Builder[T]) OrHaving(column, operator string, value any) *Builder[T] {
	b.havings.WhereList = b.havings.OrWhere(column, operator, value)
	return b
}
func (b *Builder[T]) HavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b.havings.WhereList = b.havings.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) OrHavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b.havings.WhereList = b.havings.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) HavingIn(column string, values []any) *Builder[T] {
	b.havings.WhereList = b.havings.WhereIn(column, values)
	return b
}
func (b *Builder[T]) OrHavingIn(column string, values []any) *Builder[T] {
	b.havings.WhereList = b.havings.OrWhereIn(column, values)
	return b
}
func (b *Builder[T]) HavingExists(query QueryBuilder) *Builder[T] {
	b.havings.WhereList = b.havings.WhereExists(query)
	return b
}
func (b *Builder[T]) OrHavingExists(query QueryBuilder) *Builder[T] {
	b.havings.WhereList = b.havings.OrWhereExists(query)
	return b
}
func (b *Builder[T]) HavingAnd(cb func(b *WhereList)) *Builder[T] {
	b.havings.WhereList = b.havings.And(cb)
	return b
}

func (b *Builder[T]) HavingOr(cb func(b *WhereList)) *Builder[T] {
	b.havings.WhereList = b.havings.Or(cb)
	return b
}
