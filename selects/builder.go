package selects

type Builder struct {
	selects  *selects
	from     fromTable
	wheres   *Wheres
	groupBys groupBys
	havings  *havings
	limit    *limit
	orderBys orderBys
}

func New() *Builder {
	return &Builder{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWheres(),
		groupBys: groupBys{},
		havings:  newHavings(),
		limit:    &limit{},
	}
}
