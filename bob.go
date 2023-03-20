package bob

type SelectBuilder struct {
	selects  *Selects
	from     FromTable
	wheres   *Wheres
	groupBys GroupBys
	havings  *Havings
	limit    *Limit
}

func New() *SelectBuilder {
	return NewEmpty().Select("*")
}

func From(m any) *SelectBuilder {
	return New().From(GetTable(m))
}

func NewEmpty() *SelectBuilder {
	return &SelectBuilder{
		selects:  NewSelects(),
		from:     "",
		wheres:   NewWheres(),
		groupBys: GroupBys{},
		havings:  NewHavings(),
		limit:    &Limit{},
	}
}
