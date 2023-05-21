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
	case dialects.DataTypeInt32, dialects.DataTypeUInt32, dialects.DataTypeBoolean:
		return "INTEGER"
	case dialects.DataTypeFloat32:
		return "FLOAT"
	}
	return string(t)
}

func (*SQLite) CurrentTime() string {
	return "CURRENT_TIMESTAMP"
}

func (*SQLite) Binding() string {
	return "?"
}

func UseSQLite() {
	dialects.SetDefaultDialect(func() dialects.Dialect {
		return &SQLite{}
	})
}
func init() {
	UseSQLite()
}
