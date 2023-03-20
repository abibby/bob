package bob

import "github.com/abibby/bob/dialects"

type Selects struct {
	distinct bool
	list     ExpressionList
}

func NewSelects() *Selects {
	return &Selects{
		list: ExpressionList{},
	}
}

func (s *Selects) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(s.list) == 0 {
		return "", nil, nil
	}
	r := &sqlResult{}
	r.addString("SELECT")
	if s.distinct {
		r.addString("DISTINCT")
	}
	r.add(s.list.ToSQL(d))
	return r.ToSQL(d)
}

func (b *Builder) Select(columns ...string) *Builder {
	identifiers := make([]ToSQLer, len(columns))
	for i, c := range columns {
		if c == "*" {
			identifiers[i] = Raw("*")
		} else {
			identifiers[i] = Identifier(c)
		}
	}
	b.selects.list = identifiers
	return b
}

func (b *Builder) AddSelect(columns ...string) *Builder {
	b.selects.list = append(b.selects.list, IdentifierList(columns)...)
	return b
}

func (b *Builder) SelectSubquery(sb *Builder) *Builder {
	b.selects.list = []ToSQLer{NewGroup(sb)}

	return b
}
func (b *Builder) AddSelectSubquery(sb *Builder) *Builder {
	b.selects.list = append(b.selects.list, NewGroup(sb))

	return b
}

func (b *Builder) Distinct() *Builder {
	b.selects.distinct = true
	return b
}
