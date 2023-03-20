package mysql

import (
	"strings"

	"github.com/abibby/bob/dialects"
)

type MySQL struct{}

var _ dialects.Dialect = &MySQL{}

func (*MySQL) Identifier(s string) string {
	parts := strings.Split(s, ".")
	for i, p := range parts {
		parts[i] = "`" + p + "`"
	}
	return strings.Join(parts, ".")
}
