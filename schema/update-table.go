package schema

import (
	"context"
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type UpdateTableBuilder struct {
	blueprint *Blueprint
}

var _ builder.ToSQLer = &UpdateTableBuilder{}
var _ Blueprinter = &UpdateTableBuilder{}

func Table(name string, cb func(table *Blueprint)) *UpdateTableBuilder {
	table := NewBlueprint(name)
	cb(table)
	return &UpdateTableBuilder{
		blueprint: table,
	}
}

func (b *UpdateTableBuilder) GetBlueprint() *Blueprint {
	return b.blueprint
}
func (b *UpdateTableBuilder) Type() BlueprintType {
	return BlueprintTypeUpdate
}

func (b *UpdateTableBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	alterTable := builder.Concat(builder.Raw("ALTER TABLE "), builder.Identifier(b.blueprint.name))
	for _, column := range b.blueprint.dropColumns {
		r.Add(builder.Concat(
			alterTable,
			builder.Raw(" DROP COLUMN "),
			builder.Identifier(column),
			builder.Raw(";"),
		))
	}
	for _, column := range b.blueprint.columns {
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
		b.blueprint.name,
		b.blueprint.ToGo(),
	)
}

func (b *UpdateTableBuilder) Run(ctx context.Context, tx builder.QueryExecer) error {
	return runQuery(ctx, tx, b)
}
