package selects

import (
	"fmt"
	"strings"

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
	r.Add(builder.Join(s.list, ", "))
	return r.ToSQL(d)
}

// Select sets the columns to be selected.
func (s *selects) Select(columns ...string) *selects {
	identifiers := make([]builder.ToSQLer, len(columns))
	for i, c := range columns {
		if c == "*" {
			identifiers[i] = builder.Raw("*")
		} else if strings.HasSuffix(c, ".*") {
			identifiers[i] = builder.Concat(builder.Identifier(c[:len(c)-2]), builder.Raw(".*"))
		} else {
			identifiers[i] = builder.Identifier(c)
		}
	}
	s.list = identifiers
	return s
}

// AddSelect adds new columns to be selected.
func (s *selects) AddSelect(columns ...string) *selects {
	s.list = append(s.list, builder.IdentifierList(columns)...)
	return s
}

// SelectSubquery sets a subquery to be selected.
func (s *selects) SelectSubquery(sb QueryBuilder, as string) *selects {
	return s.Select().AddSelectSubquery(sb, as)
}

// AddSelectSubquery adds a subquery to be selected.
func (s *selects) AddSelectSubquery(sb QueryBuilder, as string) *selects {
	s.list = append(s.list, builder.Concat(
		builder.Group(sb),
		builder.Raw(" as "),
		builder.Identifier(as),
	))

	return s
}

// SelectFunction sets a column to be selected with a function applied.
func (s *selects) SelectFunction(function, column string) *selects {
	return s.Select().AddSelectFunction(function, column)
}

// SelectFunction adds a column to be selected with a function applied.
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

// Distinct forces the query to only return distinct results.
func (s *selects) Distinct() *selects {
	s.distinct = true
	return s
}
