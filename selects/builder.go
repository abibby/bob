package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

type iBuilder interface {
	builder.ToSQLer
	imALittleQueryBuilderShortAndStout()
}
type Builder[T models.Model] struct {
	selects  *selects
	from     fromTable
	wheres   *Wheres
	groupBys groupBys
	havings  *havings
	limit    *limit
	orderBys orderBys
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
		wheres:   NewWheres(),
		groupBys: groupBys{},
		havings:  newHavings(),
		limit:    &limit{},
	}
}

func (*Builder[T]) imALittleQueryBuilderShortAndStout() {}
