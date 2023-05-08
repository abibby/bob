package mysql

import (
	"strings"

	"github.com/abibby/bob/dialects"
)

type MySQL struct{}

func (*MySQL) Identifier(s string) string {
	parts := strings.Split(s, ".")
	for i, p := range parts {
		parts[i] = "`" + p + "`"
	}
	return strings.Join(parts, ".")
}

func (*MySQL) DataType(t dialects.DataType) string {
	switch t {
	case dialects.DataTypeString, dialects.DataTypeJSON, dialects.DataTypeDate, dialects.DataTypeDateTime:
		return "VARCHAR(255)"
	case dialects.DataTypeInteger, dialects.DataTypeUnsignedInteger, dialects.DataTypeBoolean:
		return "INTEGER"
	case dialects.DataTypeFloat:
		return "FLOAT"
	}

	return string(t)
}

func (*MySQL) CurrentTime() string {
	return "CURRENT_TIMESTAMP"
}

func (*MySQL) TableQuery() string {
	return ""
}

func init() {
	dialects.DefaultDialect = &MySQL{}
}
