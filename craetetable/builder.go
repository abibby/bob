package craetetable

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type Builder struct {
	*Table
	// name        string
	// columns     []*ColumnBuilder
	// indexes     []*IndexBuilder
	// foreignKeys []*selects.ForeignKey
}

var _ builder.ToSQLer = &Builder{}

func CreateTable(name string, cb func(t *Table)) *Builder {
	t := &Table{
		name:        name,
		columns:     []*ColumnBuilder{},
		indexes:     []*IndexBuilder{},
		foreignKeys: []*ForeignKeyBuilder{},
	}
	cb(t)
	return &Builder{
		Table: t,
		// name:        name,
		// columns:     t.columns,
		// indexes:     t.indexes,
		// foreignKeys: t.foreignKeys,
	}
}

func (b *Builder) ToSQL(d dialects.Dialect) (string, []any, error) {
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
