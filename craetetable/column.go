package craetetable

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ColumnBuilder struct {
	name     string
	datatype dialects.DataType
	tag      *builder.Tag
	nullable bool
}

func NewColumn(name string, datatype dialects.DataType) *ColumnBuilder {
	return &ColumnBuilder{
		name:     name,
		datatype: datatype,
		tag:      &builder.Tag{},
	}
}

var _ builder.ToSQLer = &ColumnBuilder{}

func (b *ColumnBuilder) withTag(tag *builder.Tag) *ColumnBuilder {
	b.tag = tag
	return b
}
func (b *ColumnBuilder) Nullable() *ColumnBuilder {
	b.nullable = true
	return b
}
func (b *ColumnBuilder) NotNullable() *ColumnBuilder {
	b.nullable = false
	return b
}
func (b *ColumnBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.Add(builder.Identifier(b.name))
	r.AddString(d.DataType(b.datatype))
	if b.tag.Primary {
		r.AddString("PRIMARY KEY")
	}
	if b.tag.AutoIncrement {
		r.AddString("AUTOINCREMENT")
	}
	if !b.nullable {
		r.AddString("NOT NULL")
	}
	return r.ToSQL(d)
}

func (b *Builder) Columns(columns ...*ColumnBuilder) *Builder {
	b.columns = columns
	return b
}
func (b *Builder) AddColumns(columns ...*ColumnBuilder) *Builder {
	b.columns = append(b.columns, columns...)
	return b
}
