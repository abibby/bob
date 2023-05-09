package schema

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type CreateTableBuilder struct {
	blueprint   *Blueprint
	ifNotExists bool
}

var _ builder.ToSQLer = &CreateTableBuilder{}
var _ Blueprinter = &CreateTableBuilder{}

func Create(name string, cb func(b *Blueprint)) *CreateTableBuilder {
	b := NewBlueprint(name)
	cb(b)
	return &CreateTableBuilder{
		blueprint: b,
	}
}

func (b *CreateTableBuilder) GetBlueprint() *Blueprint {
	return b.blueprint
}
func (b *CreateTableBuilder) Type() BlueprintType {
	return BlueprintTypeCreate
}
func (b *CreateTableBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.AddString("CREATE TABLE")
	if b.ifNotExists {
		r.AddString("IF NOT EXISTS")
	}
	r.Add(builder.Identifier(b.blueprint.name))
	columns := make([]builder.ToSQLer, len(b.blueprint.columns))
	for i, c := range b.blueprint.columns {
		columns[i] = c
	}
	r.Add(builder.Concat(builder.Group(builder.Join(columns, ", ")), builder.Raw(";")))

	for _, index := range b.blueprint.indexes {
		r.Add(builder.Concat(index, builder.Raw(";")))
	}

	for _, foreignKey := range b.blueprint.foreignKeys {
		r.Add(builder.Concat(foreignKey, builder.Raw(";")))
	}

	return r.ToSQL(d)
}

func (b *CreateTableBuilder) ToGo() string {
	return fmt.Sprintf(
		"schema.Create(%#v, %s)",
		b.blueprint.name,
		b.blueprint.ToGo(),
	)
}
func (b *CreateTableBuilder) IfNotExists() *CreateTableBuilder {
	b.ifNotExists = true
	return b
}
