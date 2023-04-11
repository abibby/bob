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
	case dialects.DataTypeString:
		return "varchar(255)"
	case dialects.DataTypeInteger:
		return "int"
	case dialects.DataTypeFloat:
		return "float"
	}
	return string(t)
}

func init() {
	dialects.DefaultDialect = &MySQL{}
}
