package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

type QueryBuilder interface {
	builder.ToSQLer
	imALittleQueryBuilderShortAndStout()
}

//go:generate go run ../build/build.go
type SubBuilder struct {
	selects  *selects
	from     fromTable
	wheres   *WhereList
	groupBys groupBys
	havings  *WhereList
	limit    *limit
	orderBys orderBys
	scopes   scopes
}

//go:generate go run ../build/build.go
type Builder[T models.Model] struct {
	subBuilder *SubBuilder
	withs      []string
}

func New[T models.Model]() *Builder[T] {
	return NewEmpty[T]().Select("*")
}

func From[T models.Model]() *Builder[T] {
	var m T
	return New[T]().From(builder.GetTable(m))
}

func NewEmpty[T models.Model]() *Builder[T] {
	var m T
	sb := NewSubBuilder()
	sb.wheres.withParent(m)
	sb.havings.withParent(m)
	return &Builder[T]{
		subBuilder: sb,
		withs:      []string{},
	}
}
func NewSubBuilder() *SubBuilder {
	return &SubBuilder{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWhereList().withPrefix("WHERE"),
		groupBys: groupBys{},
		havings:  NewWhereList().withPrefix("HAVING"),
		limit:    &limit{},
		scopes:   scopes{},
	}
}

func (*Builder[T]) imALittleQueryBuilderShortAndStout() {}
func (*SubBuilder) imALittleQueryBuilderShortAndStout() {}
