package selects

type Builder struct {
	selects  *Selects
	from     FromTable
	wheres   *Wheres
	groupBys GroupBys
	havings  *Havings
	limit    *Limit
	orderBys OrderBys
}

func New() *Builder {
	return &Builder{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWheres(),
		groupBys: GroupBys{},
		havings:  NewHavings(),
		limit:    &Limit{},
	}
}
