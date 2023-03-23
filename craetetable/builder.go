package craetetable

import (
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

func (b *Builder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.AddString("CREATE TABLE")
	r.Add(builder.Identifier(b.name).ToSQL(d))
	r.AddString("(")
	for i, c := range b.columns {
		q, bindings, err := c.ToSQL(d)
		if i < len(b.columns)-1 {
			q += ","
		}
		r.Add(q, bindings, err)
	}
	r.AddString(")")
	return r.ToSQL(d)
}
