package bob

import "github.com/abibby/bob/dialects"

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
	q, args, err := g.sqler.ToSQL(d)
	return "(" + q + ")", args, err
}

type Literal struct{ value any }

func NewLiteral(v any) ToSQLer {
	return &Literal{value: v}
}
func (l Literal) ToSQL(d dialects.Dialect) (string, []any, error) {
	return "?", []any{l.value}, nil
}
