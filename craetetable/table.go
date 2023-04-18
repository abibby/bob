package craetetable

import (
	"github.com/abibby/bob/dialects"
)

type Table struct {
	name        string
	columns     []*ColumnBuilder
	indexes     []*IndexBuilder
	foreignKeys []*ForeignKeyBuilder
}

func (t *Table) OfType(datatype dialects.DataType, name string) *ColumnBuilder {
	c := NewColumn(name, datatype)
	t.columns = append(t.columns, c)
	return c
}
func (t *Table) String(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeString, name)
}

func (t *Table) Bool(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeBoolean, name)
}

func (t *Table) Int(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeInteger, name)
}

func (t *Table) UInt(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeUnsignedInteger, name)
}

func (t *Table) Float(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeFloat, name)
}

func (t *Table) JSON(name string) *ColumnBuilder {
	return t.OfType(dialects.DataTypeJSON, name)
}

func (t *Table) Index(name string) *IndexBuilder {
	c := &IndexBuilder{
		table: t.name,
		name:  name,
	}
	t.indexes = append(t.indexes, c)
	return c
}

func (t *Table) ForeignKey(localKey, relatedTable, relatedKey string) {
	f := &ForeignKeyBuilder{
		localKey:     localKey,
		relatedTable: relatedTable,
		relatedKey:   relatedKey,
	}
	t.foreignKeys = append(t.foreignKeys, f)
	return
}
