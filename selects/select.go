package selects

import (
	"fmt"

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

func (s *selects) Clone() *selects {
	return &selects{
		distinct: s.distinct,
		list:     cloneSlice(s.list),
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

func (s *selects) Select(columns ...string) *selects {
	identifiers := make([]builder.ToSQLer, len(columns))
	for i, c := range columns {
		if c == "*" {
			identifiers[i] = builder.Raw("*")
		} else {
			identifiers[i] = builder.Identifier(c)
		}
	}
	s.list = identifiers
	return s
}

func (s *selects) AddSelect(columns ...string) *selects {
	s.list = append(s.list, builder.IdentifierList(columns)...)
	return s
}

func (s *selects) SelectSubquery(sb QueryBuilder, as string) *selects {
	s.list = make([]builder.ToSQLer, 0, 1)
	return s.AddSelectSubquery(sb, as)
}
func (s *selects) AddSelectSubquery(sb QueryBuilder, as string) *selects {
	s.list = append(s.list, builder.Concat(
		builder.Group(sb),
		builder.Raw(" as "),
		builder.Identifier(as),
	))

	return s
}
func (s *selects) SelectFunction(function, column string) *selects {
	return s.Select().AddSelectFunction(function, column)
}
func (s *selects) AddSelectFunction(function, column string) *selects {
	s.list = append(s.list, builder.ToSQLFunc(func(d dialects.Dialect) (string, []any, error) {
		var c builder.ToSQLer
		if column == "*" {
			c = builder.Raw("*")
		} else {
			c = builder.Identifier(column)
		}
		q, bindings, err := c.ToSQL(d)
		if err != nil {
			return "", nil, err
		}
		return fmt.Sprintf("%s(%s)", function, q), bindings, nil
	}))

	return s
}

func (s *selects) Distinct() *selects {
	s.distinct = true
	return s
}
