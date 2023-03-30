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
type Builder[T models.Model] struct {
	selects  *selects
	from     fromTable
	wheres   *WhereList
	groupBys groupBys
	havings  *WhereList
	limit    *limit
	orderBys orderBys

	withs []string
}

func New[T models.Model]() *Builder[T] {
	return NewEmpty[T]().Select("*")
}

func From[T models.Model]() *Builder[T] {
	var m T
	return New[T]().From(builder.GetTable(m))
}

func NewEmpty[T models.Model]() *Builder[T] {
	return &Builder[T]{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWhereList().withPrefix("WHERE"),
		groupBys: groupBys{},
		havings:  NewWhereList().withPrefix("HAVING"),
		limit:    &limit{},
	}
}

func (*Builder[T]) imALittleQueryBuilderShortAndStout() {}
