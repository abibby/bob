package selects

import "context"

func (b *Builder[T]) WithContext(ctx context.Context) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WithContext(ctx)
	return b
}
func (b *Builder[T]) From(table string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.From(table)
	return b
}
func (b *Builder[T]) GroupBy(columns ...string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.GroupBy(columns...)
	return b
}
func (b *Builder[T]) AddGroupBy(columns ...string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.AddGroupBy(columns...)
	return b
}
func (b *Builder[T]) Limit(limit int) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Limit(limit)
	return b
}
func (b *Builder[T]) Offset(offset int) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Offset(offset)
	return b
}
func (b *Builder[T]) OrderBy(column string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrderBy(column)
	return b
}
func (b *Builder[T]) OrderByDesc(column string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrderByDesc(column)
	return b
}
func (b *Builder[T]) WithScope(scope *Scope) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WithScope(scope)
	return b
}
func (b *Builder[T]) WithoutScope(scope *Scope) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WithoutScope(scope)
	return b
}
func (b *Builder[T]) WithoutGlobalScope(scope *Scope) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WithoutGlobalScope(scope)
	return b
}
func (b *Builder[T]) Select(columns ...string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Select(columns...)
	return b
}
func (b *Builder[T]) AddSelect(columns ...string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.AddSelect(columns...)
	return b
}
func (b *Builder[T]) SelectSubquery(sb QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.SelectSubquery(sb)
	return b
}
func (b *Builder[T]) AddSelectSubquery(sb QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.AddSelectSubquery(sb)
	return b
}
func (b *Builder[T]) SelectFunction(function, column string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.SelectFunction(function, column)
	return b
}
func (b *Builder[T]) AddSelectFunction(function, column string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.AddSelectFunction(function, column)
	return b
}
func (b *Builder[T]) Distinct() *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Distinct()
	return b
}
func (b *Builder[T]) Where(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Where(column, operator, value)
	return b
}
func (b *Builder[T]) Having(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Having(column, operator, value)
	return b
}
func (b *Builder[T]) OrWhere(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhere(column, operator, value)
	return b
}
func (b *Builder[T]) OrHaving(column, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHaving(column, operator, value)
	return b
}
func (b *Builder[T]) WhereColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) HavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) OrWhereColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) OrHavingColumn(column, operator string, valueColumn string) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingColumn(column, operator, valueColumn)
	return b
}
func (b *Builder[T]) WhereIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereIn(column, values)
	return b
}
func (b *Builder[T]) HavingIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingIn(column, values)
	return b
}
func (b *Builder[T]) OrWhereIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereIn(column, values)
	return b
}
func (b *Builder[T]) OrHavingIn(column string, values []any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingIn(column, values)
	return b
}
func (b *Builder[T]) WhereExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereExists(query)
	return b
}
func (b *Builder[T]) HavingExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingExists(query)
	return b
}
func (b *Builder[T]) OrWhereExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereExists(query)
	return b
}
func (b *Builder[T]) OrHavingExists(query QueryBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingExists(query)
	return b
}
func (b *Builder[T]) WhereSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) HavingSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) OrWhereSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) OrHavingSubquery(subquery QueryBuilder, operator string, value any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingSubquery(subquery, operator, value)
	return b
}
func (b *Builder[T]) WhereHas(relation string, cb func(q *SubBuilder) *SubBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereHas(relation, cb)
	return b
}
func (b *Builder[T]) HavingHas(relation string, cb func(q *SubBuilder) *SubBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingHas(relation, cb)
	return b
}
func (b *Builder[T]) OrWhereHas(relation string, cb func(q *SubBuilder) *SubBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereHas(relation, cb)
	return b
}
func (b *Builder[T]) OrHavingHas(relation string, cb func(q *SubBuilder) *SubBuilder) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingHas(relation, cb)
	return b
}
func (b *Builder[T]) WhereRaw(rawSql string, bindings ...any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.WhereRaw(rawSql, bindings...)
	return b
}
func (b *Builder[T]) HavingRaw(rawSql string, bindings ...any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingRaw(rawSql, bindings...)
	return b
}
func (b *Builder[T]) OrWhereRaw(rawSql string, bindings ...any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrWhereRaw(rawSql, bindings...)
	return b
}
func (b *Builder[T]) OrHavingRaw(rawSql string, bindings ...any) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.OrHavingRaw(rawSql, bindings...)
	return b
}
func (b *Builder[T]) And(cb func(wl *WhereList)) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.And(cb)
	return b
}
func (b *Builder[T]) HavingAnd(cb func(wl *WhereList)) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingAnd(cb)
	return b
}
func (b *Builder[T]) Or(cb func(wl *WhereList)) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Or(cb)
	return b
}
func (b *Builder[T]) HavingOr(cb func(wl *WhereList)) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.HavingOr(cb)
	return b
}
func (b *Builder[T]) Dump(ctx context.Context) *Builder[T] {
	b = b.Clone()
	b.subBuilder = b.subBuilder.Dump(ctx)
	return b
}
