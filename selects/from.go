package selects

import (
	"github.com/abibby/bob/dialects"
)

type fromTable string

func (f fromTable) ToSQL(d dialects.Dialect) (string, []any, error) {
	if f == "" {
		return "", nil, nil
	}

	return "FROM " + d.Identifier(string(f)), nil, nil
}

func (b *Builder) From(table string) *Builder {
	b.from = fromTable(table)
	return b
}
