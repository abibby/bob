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
	case dialects.DataTypeString, dialects.DataTypeJSON, dialects.DataTypeDate, dialects.DataTypeDateTime:
		return "TEXT"
	case dialects.DataTypeInteger, dialects.DataTypeUnsignedInteger, dialects.DataTypeBoolean:
		return "INTEGER"
	case dialects.DataTypeFloat:
		return "FLOAT"
	}
	return string(t)
}

func init() {
	dialects.DefaultDialect = &SQLite{}
}
