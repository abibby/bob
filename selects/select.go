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

func (b *Builder[T]) Select(columns ...string) *Builder[T] {
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

func (b *Builder[T]) AddSelect(columns ...string) *Builder[T] {
	b.selects.list = append(b.selects.list, builder.IdentifierList(columns)...)
	return b
}

func (b *Builder[T]) SelectSubquery(sb *Builder[T]) *Builder[T] {
	b.selects.list = []builder.ToSQLer{builder.NewGroup(sb)}

	return b
}
func (b *Builder[T]) AddSelectSubquery(sb *Builder[T]) *Builder[T] {
	b.selects.list = append(b.selects.list, builder.NewGroup(sb))

	return b
}
func (b *Builder[T]) SelectFunction(function, column string) *Builder[T] {
	b.selects.list = append(b.selects.list, builder.ToSQLFunc(func(d dialects.Dialect) (string, []any, error) {
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

	return b
}

func (b *Builder[T]) Distinct() *Builder[T] {
	b.selects.distinct = true
	return b
}
