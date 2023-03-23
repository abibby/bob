package craetetable

import "github.com/abibby/bob/dialects"

type Table struct {
	columns []*ColumnBuilder
}

func (t *Table) String(name string) *ColumnBuilder {
	c := &ColumnBuilder{
		name:     name,
		datatype: dialects.DataTypeString,
	}
	t.columns = append(t.columns, c)
	return c
}

func (t *Table) Int(name string) *ColumnBuilder {
	c := &ColumnBuilder{
		name:     name,
		datatype: dialects.DataTypeInteger,
	}
	t.columns = append(t.columns, c)
	return c
}

func (t *Table) Float(name string) *ColumnBuilder {
	c := &ColumnBuilder{
		name:     name,
		datatype: dialects.DataTypeFloat,
	}
	t.columns = append(t.columns, c)
	return c
}
