package schema

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type UpdateTableBuilder struct {
	*Blueprint
}

var _ builder.ToSQLer = &UpdateTableBuilder{}

func Table(name string, cb func(table *Blueprint)) *UpdateTableBuilder {
	table := newBlueprint(name)
	cb(table)
	return &UpdateTableBuilder{
		Blueprint: table,
	}
}

func (b *UpdateTableBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	return "", nil, nil
}

func (b *UpdateTableBuilder) ToGo() string {
	return fmt.Sprintf(
		"schema.Table(%#v, %s)",
		b.name,
		b.Blueprint.ToGo(),
	)
}
