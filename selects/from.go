package selects

import (
	"context"

	"github.com/abibby/bob/dialects"
)

type fromTable string

func (f fromTable) Clone() fromTable {
	return f
}
func (f fromTable) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	if f == "" {
		return "", nil, nil
	}

	return "FROM " + d.Identifier(string(f)), nil, nil
}

func (f fromTable) From(table string) fromTable {
	return fromTable(table)
}
