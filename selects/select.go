package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type selects struct {
	distinct bool
	list     []builder.ToSQLer
}

func NewSelects() *selects {
	return &selects{
		list: []builder.ToSQLer{},
	}
}

func (s *selects) ToSQL(d dialects.Dialect) (string, []any, error) {
	if len(s.list) == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("SELECT")
	if s.distinct {
		r.AddString("DISTINCT")
	}
	r.Add(builder.Join(s.list, ", ").ToSQL(d))
	return r.ToSQL(d)
}

func (b *Builder) Select(columns ...string) *Builder {
	identifiers := make([]builder.ToSQLer, len(columns))
	for i, c := range columns {
		if c == "*" {
			identifiers[i] = builder.Raw("*")
		} else {
			identifiers[i] = builder.Identifier(c)
		}
	}
	b.selects.list = identifiers
	return b
}

func (b *Builder) AddSelect(columns ...string) *Builder {
	b.selects.list = append(b.selects.list, builder.IdentifierList(columns)...)
	return b
}

func (b *Builder) SelectSubquery(sb *Builder) *Builder {
	b.selects.list = []builder.ToSQLer{builder.NewGroup(sb)}

	return b
}
func (b *Builder) AddSelectSubquery(sb *Builder) *Builder {
	b.selects.list = append(b.selects.list, builder.NewGroup(sb))

	return b
}

func (b *Builder) Distinct() *Builder {
	b.selects.distinct = true
	return b
}