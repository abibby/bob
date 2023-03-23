package builder

import (
	"github.com/abibby/bob/dialects"
)

type ToSQLer interface {
	ToSQL(d dialects.Dialect) (string, []any, error)
}

type toSQLFunc func(d dialects.Dialect) (string, []any, error)

func (f toSQLFunc) ToSQL(d dialects.Dialect) (string, []any, error) {
	return f(d)
}

type Identifier string

func (i Identifier) ToSQL(d dialects.Dialect) (string, []any, error) {
	return d.Identifier(string(i)), nil, nil
}

func IdentifierList(strs []string) []ToSQLer {
	identifiers := make([]ToSQLer, len(strs))
	for i, s := range strs {
		identifiers[i] = Identifier(s)
	}
	return identifiers
}

func Join(sqlers []ToSQLer, sep string) ToSQLer {
	return toSQLFunc(func(d dialects.Dialect) (string, []any, error) {
		sql := ""
		bindings := []any{}
		for i, sqler := range sqlers {
			q, b, err := sqler.ToSQL(d)
			if err != nil {
				return "", nil, err
			}
			sql += q
			if i < len(sqlers)-1 {
				sql += sep
			}
			bindings = append(bindings, b...)
		}
		return sql, bindings, nil
	})
}

type Raw string

func (r Raw) ToSQL(d dialects.Dialect) (string, []any, error) {
	return string(r), nil, nil
}

type Group struct {
	sqler ToSQLer
}

func NewGroup(sqler ToSQLer) ToSQLer {
	return &Group{sqler: sqler}
}

func (g *Group) ToSQL(d dialects.Dialect) (string, []any, error) {
	q, bindings, err := g.sqler.ToSQL(d)
	return "(" + q + ")", bindings, err
}

type Literal struct{ value any }

func NewLiteral(v any) ToSQLer {
	return &Literal{value: v}
}
func (l Literal) ToSQL(d dialects.Dialect) (string, []any, error) {
	return "?", []any{l.value}, nil
}

func LiteralList(values []any) []ToSQLer {
	literals := make([]ToSQLer, len(values))
	for i, s := range values {
		literals[i] = NewLiteral(s)
	}
	return literals
}
