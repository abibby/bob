package schema

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ColumnBuilder struct {
	name     string
	datatype dialects.DataType

	nullable           bool
	primary            bool
	autoIncrement      bool
	change             bool
	defaultValue       builder.ToSQLer
	afterColumn        string
	unique             bool
	defaultCurrentTime bool
	index              bool
}

func NewColumn(name string, datatype dialects.DataType) *ColumnBuilder {
	return &ColumnBuilder{
		name:     name,
		datatype: datatype,
	}
}

var _ builder.ToSQLer = &ColumnBuilder{}

func (b *ColumnBuilder) Equals(newB *ColumnBuilder) bool {
	return b.datatype == newB.datatype &&
		b.nullable == newB.nullable &&
		b.autoIncrement == newB.autoIncrement &&
		b.index == newB.index &&
		b.name == newB.name &&
		b.primary == newB.primary
}

func (b *ColumnBuilder) Name() string {
	return b.name
}

func (b *ColumnBuilder) Nullable() *ColumnBuilder {
	b.nullable = true
	return b
}
func (b *ColumnBuilder) NotNullable() *ColumnBuilder {
	b.nullable = false
	return b
}
func (b *ColumnBuilder) Primary() *ColumnBuilder {
	b.primary = true
	return b
}
func (b *ColumnBuilder) AutoIncrement() *ColumnBuilder {
	b.autoIncrement = true
	return b
}
func (b *ColumnBuilder) After(column string) *ColumnBuilder {
	b.afterColumn = column
	return b
}
func (b *ColumnBuilder) Change() *ColumnBuilder {
	b.change = true
	return b
}
func (b *ColumnBuilder) Default(v any) *ColumnBuilder {
	b.defaultValue = builder.Literal(v)
	return b
}
func (b *ColumnBuilder) Type(datatype dialects.DataType) *ColumnBuilder {
	b.datatype = datatype
	return b
}
func (b *ColumnBuilder) Unique() *ColumnBuilder {
	b.unique = true
	return b
}
func (b *ColumnBuilder) DefaultCurrentTime() *ColumnBuilder {
	b.defaultCurrentTime = true
	return b
}
func (b *ColumnBuilder) Index() *ColumnBuilder {
	b.index = true
	return b
}
func (b *ColumnBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.Add(builder.Identifier(b.name))
	r.AddString(d.DataType(b.datatype))
	if b.primary {
		r.AddString("PRIMARY KEY")
	}
	if b.autoIncrement {
		r.AddString("AUTOINCREMENT")
	}
	if !b.nullable {
		r.AddString("NOT NULL")
	}
	if b.defaultValue != nil {
		r.AddString("DEFAULT").
			Add(b.defaultValue)
	}
	if b.defaultCurrentTime {
		r.AddString("DEFAULT").
			AddString(d.CurrentTime())
	}
	if b.unique {
		r.AddString("UNIQUE")
	}
	return r.ToSQL(d)
}

func (b *ColumnBuilder) ToGo() string {
	src := ""
	if b.primary {
		src += ".Primary()"
	}
	if b.autoIncrement {
		src += ".AutoIncrement()"
	}
	if b.nullable {
		src += ".Nullable()"
	}
	if b.defaultValue != nil {
		src += fmt.Sprintf(".Default(%#v)", b.defaultValue)
	}
	if b.defaultCurrentTime {
		src += ".DefaultCurrentTime()"
	}
	if b.unique {
		src += ".Unique()"
	}
	if b.index {
		src += ".Index()"
	}
	if b.change {
		src += ".Change()"
	}
	return src
}

func (b *CreateTableBuilder) Columns(columns ...*ColumnBuilder) *CreateTableBuilder {
	b.columns = columns
	return b
}
func (b *CreateTableBuilder) AddColumns(columns ...*ColumnBuilder) *CreateTableBuilder {
	b.columns = append(b.columns, columns...)
	return b
}
