package bob

import (
	"github.com/abibby/bob/dialects"
)

type FromTable string

func (f FromTable) ToSQL(d dialects.Dialect) (string, []any, error) {
	if f == "" {
		return "", nil, nil
	}

	return "FROM " + d.Identifier(string(f)), nil, nil
}

func (b *Builder) From(table string) *Builder {
	b.from = FromTable(table)
	return b
}
