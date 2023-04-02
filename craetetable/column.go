package craetetable

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ColumnBuilder struct {
	name     string
	datatype dialects.DataType
}

func NewColumn(name string) *ColumnBuilder {
	return &ColumnBuilder{}
}

var _ builder.ToSQLer = &ColumnBuilder{}

func (b *ColumnBuilder) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.Add(builder.Identifier(b.name).ToSQL(ctx, d))
	r.AddString(d.DataType(b.datatype))
	return r.ToSQL(ctx, d)
}

func (b *Builder) Columns(columns ...*ColumnBuilder) *Builder {
	b.columns = columns
	return b
}
func (b *Builder) AddColumns(columns ...*ColumnBuilder) *Builder {
	b.columns = append(b.columns, columns...)
	return b
}
