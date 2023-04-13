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

func (f fromTable) From(table string) fromTable {
	return fromTable(table)
}

func (b *SubBuilder) GetTable() string {
	return string(b.from)
}
func (b *Builder[T]) GetTable() string {
	return b.subBuilder.GetTable()
}
