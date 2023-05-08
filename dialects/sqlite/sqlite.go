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
	case dialects.DataTypeString, dialects.DataTypeJSON:
		return "TEXT"
	case dialects.DataTypeDate, dialects.DataTypeDateTime:
		return "TIMESTAMP"
	case dialects.DataTypeInteger, dialects.DataTypeUnsignedInteger, dialects.DataTypeBoolean:
		return "INTEGER"
	case dialects.DataTypeFloat:
		return "FLOAT"
	}
	return string(t)
}

func (*SQLite) CurrentTime() string {
	return "CURRENT_TIMESTAMP"
}

func (*SQLite) TableQuery() string {
	return ""
}

func init() {
	dialects.DefaultDialect = &SQLite{}
}
