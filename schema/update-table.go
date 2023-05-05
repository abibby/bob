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
var _ Blueprinter = &UpdateTableBuilder{}

func Table(name string, cb func(table *Blueprint)) *UpdateTableBuilder {
	table := newBlueprint(name)
	cb(table)
	return &UpdateTableBuilder{
		Blueprint: table,
	}
}
func (b *UpdateTableBuilder) Type() BlueprintType {
	return BlueprintTypeUpdate
}

func (b *UpdateTableBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	alterTable := builder.Concat(builder.Raw("ALTER TABLE "), builder.Identifier(b.name))
	for _, column := range b.dropColumns {
		r.Add(builder.Concat(
			alterTable,
			builder.Raw(" DROP COLUMN "),
			builder.Identifier(column),
			builder.Raw(";"),
		))
	}
	for _, column := range b.columns {
		if column.change {
			r.Add(builder.Concat(
				alterTable,
				builder.Raw(" MODIFY COLUMN "),
				column,
				builder.Raw(";"),
			))
		} else {
			r.Add(builder.Concat(
				alterTable,
				builder.Raw(" ADD "),
				column,
				builder.Raw(";"),
			))
		}
	}
	return r.ToSQL(d)
}

func (b *UpdateTableBuilder) ToGo() string {
	return fmt.Sprintf(
		"schema.Table(%#v, %s)",
		b.name,
		b.Blueprint.ToGo(),
	)
}
