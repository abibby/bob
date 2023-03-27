package selects

import (
	"github.com/abibby/bob/dialects"
)

type fromTable string

func (f fromTable) Clone() fromTable {
	return f
}
func (f fromTable) ToSQL(d dialects.Dialect) (string, []any, error) {
	if f == "" {
		return "", nil, nil
	}

	return "FROM " + d.Identifier(string(f)), nil, nil
}

func (b *Builder[T]) From(table string) *Builder[T] {
	b.from = fromTable(table)
	return b
}
