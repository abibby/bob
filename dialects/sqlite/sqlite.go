package sqlite

import (
	"strings"

	"github.com/abibby/bob/dialects"
)

type SQLite struct{}

func (*SQLite) Identifier(s string) string {
	parts := strings.Split(s, ".")
	for i, p := range parts {
		parts[i] = "\"" + p + "\""
	}
	return strings.Join(parts, ".")
}

func (*SQLite) DataType(t dialects.DataType) string {
	switch t {
	case dialects.DataTypeString:
		return "text"
	case dialects.DataTypeInteger:
		return "int"
	case dialects.DataTypeFloat:
		return "float"
	}
	return string(t)
}

func (d *SQLite) InsertOrUpdate(table, primaryKey string, columns []string, values []any) (string, []any) {

	return "", nil
}

func init() {
	dialects.DefaultDialect = &SQLite{}
}
