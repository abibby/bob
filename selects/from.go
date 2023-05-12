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

// From sets the table which the query is targeting.
func (f fromTable) From(table string) fromTable {
	return fromTable(table)
}

// GetTable returns the table the query is targeting
func (b *SubBuilder) GetTable() string {
	return string(b.from)
}

// GetTable returns the table the query is targeting
func (b *Builder[T]) GetTable() string {
	return b.subBuilder.GetTable()
}
