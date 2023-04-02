package craetetable

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type Builder struct {
	name    string
	columns []*ColumnBuilder
}

var _ builder.ToSQLer = &Builder{}

func CreateTable(name string, cb func(t *Table)) *Builder {
	t := &Table{
		columns: []*ColumnBuilder{},
	}
	cb(t)
	return &Builder{
		name:    name,
		columns: t.columns,
	}
}

func (b *Builder) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.AddString("CREATE TABLE")
	r.Add(builder.Identifier(b.name).ToSQL(ctx, d))
	columns := make([]builder.ToSQLer, len(b.columns))
	for i, c := range b.columns {
		columns[i] = c
	}
	r.Add(builder.Group(builder.Join(columns, ", ")).ToSQL(ctx, d))
	// r.AddString("(")
	// for i, c := range b.columns {
	// 	q, bindings, err := c.ToSQL(ctx, d)
	// 	if i < len(b.columns)-1 {
	// 		q += ","
	// 	}
	// 	r.Add(q, bindings, err)
	// }
	// r.AddString(")")
	return r.ToSQL(ctx, d)
}
