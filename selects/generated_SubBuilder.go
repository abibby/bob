package selects

func (b *SubBuilder) From(table string) *SubBuilder {
	b = b.Clone()
	b.from = b.from.From(table)
	return b
}
func (b *SubBuilder) GroupBy(columns ...string) *SubBuilder {
	b = b.Clone()
	b.groupBys = b.groupBys.GroupBy(columns...)
	return b
}
func (b *SubBuilder) AddGroupBy(columns ...string) *SubBuilder {
	b = b.Clone()
	b.groupBys = b.groupBys.AddGroupBy(columns...)
	return b
}
func (b *SubBuilder) Limit(limit int) *SubBuilder {
	b = b.Clone()
	b.limit = b.limit.Limit(limit)
	return b
}
func (b *SubBuilder) Offset(offset int) *SubBuilder {
	b = b.Clone()
	b.limit = b.limit.Offset(offset)
	return b
}
func (b *SubBuilder) OrderBy(column string) *SubBuilder {
	b = b.Clone()
	b.orderBys = b.orderBys.OrderBy(column)
	return b
}
func (b *SubBuilder) OrderByDesc(column string) *SubBuilder {
	b = b.Clone()
	b.orderBys = b.orderBys.OrderByDesc(column)
	return b
}
func (b *SubBuilder) Select(columns ...string) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.Select(columns...)
	return b
}
func (b *SubBuilder) AddSelect(columns ...string) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.AddSelect(columns...)
	return b
}
func (b *SubBuilder) SelectSubquery(sb QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.SelectSubquery(sb)
	return b
}
func (b *SubBuilder) AddSelectSubquery(sb QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.AddSelectSubquery(sb)
	return b
}
func (b *SubBuilder) SelectFunction(function, column string) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.SelectFunction(function, column)
	return b
}
func (b *SubBuilder) AddSelectFunction(function, column string) *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.AddSelectFunction(function, column)
	return b
}
func (b *SubBuilder) Distinct() *SubBuilder {
	b = b.Clone()
	b.selects = b.selects.Distinct()
	return b
}
func (b *SubBuilder) Where(column, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.Where(column, operator, value)
	return b
}
func (b *SubBuilder) Having(column, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.Where(column, operator, value)
	return b
}
func (b *SubBuilder) OrWhere(column, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhere(column, operator, value)
	return b
}
func (b *SubBuilder) OrHaving(column, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhere(column, operator, value)
	return b
}
func (b *SubBuilder) WhereColumn(column, operator string, valueColumn string) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *SubBuilder) HavingColumn(column, operator string, valueColumn string) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *SubBuilder) OrWhereColumn(column, operator string, valueColumn string) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *SubBuilder) OrHavingColumn(column, operator string, valueColumn string) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *SubBuilder) WhereIn(column string, values []any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.WhereIn(column, values)
	return b
}
func (b *SubBuilder) HavingIn(column string, values []any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.WhereIn(column, values)
	return b
}
func (b *SubBuilder) OrWhereIn(column string, values []any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereIn(column, values)
	return b
}
func (b *SubBuilder) OrHavingIn(column string, values []any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhereIn(column, values)
	return b
}
func (b *SubBuilder) WhereExists(query QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.WhereExists(query)
	return b
}
func (b *SubBuilder) HavingExists(query QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.WhereExists(query)
	return b
}
func (b *SubBuilder) OrWhereExists(query QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereExists(query)
	return b
}
func (b *SubBuilder) OrHavingExists(query QueryBuilder) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhereExists(query)
	return b
}
func (b *SubBuilder) WhereSubquery(subquery QueryBuilder, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.WhereSubquery(subquery, operator, value)
	return b
}
func (b *SubBuilder) HavingSubquery(subquery QueryBuilder, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.WhereSubquery(subquery, operator, value)
	return b
}
func (b *SubBuilder) OrWhereSubquery(subquery QueryBuilder, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereSubquery(subquery, operator, value)
	return b
}
func (b *SubBuilder) OrHavingSubquery(subquery QueryBuilder, operator string, value any) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhereSubquery(subquery, operator, value)
	return b
}
func (b *SubBuilder) WhereHas(relation string, cb func(q *SubBuilder) *SubBuilder) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.WhereHas(relation, cb)
	return b
}
func (b *SubBuilder) HavingHas(relation string, cb func(q *SubBuilder) *SubBuilder) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.WhereHas(relation, cb)
	return b
}
func (b *SubBuilder) OrWhereHas(relation string, cb func(q *SubBuilder) *SubBuilder) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.OrWhereHas(relation, cb)
	return b
}
func (b *SubBuilder) OrHavingHas(relation string, cb func(q *SubBuilder) *SubBuilder) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.OrWhereHas(relation, cb)
	return b
}
func (b *SubBuilder) And(cb func(wl *WhereList)) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.And(cb)
	return b
}
func (b *SubBuilder) HavingAnd(cb func(wl *WhereList)) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.And(cb)
	return b
}
func (b *SubBuilder) Or(cb func(wl *WhereList)) *SubBuilder {
	b = b.Clone()
	b.wheres = b.wheres.Or(cb)
	return b
}
func (b *SubBuilder) HavingOr(cb func(wl *WhereList)) *SubBuilder {
	b = b.Clone()
	b.havings = b.havings.Or(cb)
	return b
}
