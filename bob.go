package bob

type Builder struct {
	selects  *Selects
	from     FromTable
	wheres   *Wheres
	groupBys GroupBys
	havings  *Havings
}

func New() *Builder {
	return NewEmpty().Select("*")
}

func From(m any) *Builder {
	return New().From(GetTable(m))
}

func NewEmpty() *Builder {
	return &Builder{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWheres(),
		groupBys: GroupBys{},
		havings:  NewHavings(),
	}
}
