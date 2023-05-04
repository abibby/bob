package schema

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type CreateTableBuilder struct {
	*Blueprint
}

var _ builder.ToSQLer = &CreateTableBuilder{}

func Create(name string, cb func(b *Blueprint)) *CreateTableBuilder {
	b := newBlueprint(name)
	cb(b)
	return &CreateTableBuilder{
		Blueprint: b,
	}
}

func (b *CreateTableBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.AddString("CREATE TABLE")
	r.Add(builder.Identifier(b.name))
	columns := make([]builder.ToSQLer, len(b.columns))
	for i, c := range b.columns {
		columns[i] = c
	}
	r.Add(builder.Concat(builder.Group(builder.Join(columns, ", ")), builder.Raw(";")))

	for _, index := range b.indexes {
		r.Add(builder.Concat(index, builder.Raw(";")))
	}

	for _, foreignKey := range b.foreignKeys {
		r.Add(builder.Concat(foreignKey, builder.Raw(";")))
	}

	return r.ToSQL(d)
}

func (b *CreateTableBuilder) ToGo() string {
	return fmt.Sprintf(
		"schema.Create(%#v, %s)",
		b.name,
		b.Blueprint.ToGo(),
	)
}
