package schema

import (
	"fmt"

	"github.com/abibby/bob/dialects"
)

type Blueprint struct {
	name        string
	columns     []*ColumnBuilder
	dropColumns []string
	indexes     []*IndexBuilder
	foreignKeys []*ForeignKeyBuilder
}

func newBlueprint(name string) *Blueprint {
	return &Blueprint{
		name:        name,
		columns:     []*ColumnBuilder{},
		dropColumns: []string{},
		indexes:     []*IndexBuilder{},
		foreignKeys: []*ForeignKeyBuilder{},
	}
}

func (t *Blueprint) OfType(datatype dialects.DataType, name string) *ColumnBuilder {
	c := NewColumn(name, datatype)
	t.columns = append(t.columns, c)
	return c
}
func (t *Blueprint) String(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeString, name)
}

func (t *Blueprint) Bool(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeBoolean, name)
}

func (t *Blueprint) Int(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeInteger, name)
}

func (t *Blueprint) UInt(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeUnsignedInteger, name)
}

func (t *Blueprint) Float(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeFloat, name)
}

func (t *Blueprint) JSON(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeJSON, name)
}
func (t *Blueprint) Date(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeDate, name)
}
func (t *Blueprint) DateTime(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeDateTime, name)
}

func (t *Blueprint) Index(name string) *IndexBuilder {
	c := &IndexBuilder{
		table: t.name,
		name:  name,
	}
	t.indexes = append(t.indexes, c)
	return c
}

func (t *Blueprint) ForeignKey(localKey, relatedTable, relatedKey string) {
	f := &ForeignKeyBuilder{
		localKey:     localKey,
		relatedTable: relatedTable,
		relatedKey:   relatedKey,
	}
	t.foreignKeys = append(t.foreignKeys, f)
}

func (t *Blueprint) DropColumn(column string) {
	t.dropColumns = append(t.dropColumns, column)
}

func (b *Blueprint) ToGo() string {
	src := "func(table *schema.Blueprint) {\n"
	for _, c := range b.columns {
		m := map[dialects.DataType]string{
			dialects.DataTypeString:          "String",
			dialects.DataTypeInteger:         "Int",
			dialects.DataTypeUnsignedInteger: "UInt",
			dialects.DataTypeFloat:           "Float",
			dialects.DataTypeBoolean:         "Boolean",
			dialects.DataTypeJSON:            "JSON",
			dialects.DataTypeDate:            "Date",
			dialects.DataTypeDateTime:        "DateTime",
		}
		src += fmt.Sprintf("\ttable.%s(%#v)%s\n", m[c.datatype], c.name, c.ToGo())
	}

	return src + "}"
}
