package selects

func (b *Builder[T]) AddSelect(columns ...string) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.AddSelect(columns...)
	return b
}
func (b *Builder[T]) AddSelectFunction(function, column string) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.AddSelectFunction(function, column)
	return b
}
func (b *Builder[T]) AddSelectSubquery(sb QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.AddSelectSubquery(sb)
	return b
}
func (b *Builder[T]) Distinct() *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.Distinct()
	return b
}
func (b *Builder[T]) Select(columns ...string) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.Select(columns...)
	return b
}
func (b *Builder[T]) SelectFunction(function, column string) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.SelectFunction(function, column)
	return b
}
func (b *Builder[T]) SelectSubquery(sb QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.selects = b.selects.SelectSubquery(sb)
	return b
}
func (b *Builder[T]) From(table string) *Builder[T] {
	b = b.Clone()
	b.from = b.from.From(table)
	return b
}
func (b *Builder[T]) And(wl *WhereList) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.And(wl)
	return b
}
func (b *Builder[T]) Or(wl *WhereList) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.Or(wl)
	return b
}
func (b *Builder[T]) OrWhere(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.OrWhere(column, operator, value)
	return b
}
func (b *Builder[T]) OrWhereColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) OrWhereExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereExists(query)
	return b
}
func (b *Builder[T]) OrWhereIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereIn(column, values)
	return b
}
func (b *Builder[T]) OrWhereSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) Where(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.Where(column, operator, value)
	return b
}
func (b *Builder[T]) WhereColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) WhereExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.WhereExists(query)
	return b
}
func (b *Builder[T]) WhereIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.WhereIn(column, values)
	return b
}
func (b *Builder[T]) WhereSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.wheres = b.wheres.WhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) AddGroupBy(columns ...string) *Builder[T] {
	b = b.Clone()
	b.groupBys = b.groupBys.AddGroupBy(columns...)
	return b
}
func (b *Builder[T]) GroupBy(columns ...string) *Builder[T] {
	b = b.Clone()
	b.groupBys = b.groupBys.GroupBy(columns...)
	return b
}
func (b *Builder[T]) HavingAnd(wl *WhereList) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.And(wl)
	return b
}
func (b *Builder[T]) HavingOr(wl *WhereList) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.Or(wl)
	return b
}
func (b *Builder[T]) OrHaving(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.OrWhere(column, operator, value)
	return b
}
func (b *Builder[T]) OrHavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) OrHavingExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.OrWhereExists(query)
	return b
}
func (b *Builder[T]) OrHavingIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.OrWhereIn(column, values)
	return b
}
func (b *Builder[T]) OrHavingSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.OrWhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) Having(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.Where(column, operator, value)
	return b
}
func (b *Builder[T]) HavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) HavingExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.WhereExists(query)
	return b
}
func (b *Builder[T]) HavingIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.WhereIn(column, values)
	return b
}
func (b *Builder[T]) HavingSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.havings = b.havings.WhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) Limit(limit int) *Builder[T] {
	b = b.Clone()
	b.limit = b.limit.Limit(limit)
	return b
}
func (b *Builder[T]) Offset(offset int) *Builder[T] {
	b = b.Clone()
	b.limit = b.limit.Offset(offset)
	return b
}
func (b *Builder[T]) OrderBy(column string) *Builder[T] {
	b = b.Clone()
	b.orderBys = b.orderBys.OrderBy(column)
	return b
}
func (b *Builder[T]) OrderByDesc(column string) *Builder[T] {
	b = b.Clone()
	b.orderBys = b.orderBys.OrderByDesc(column)
	return b
}
