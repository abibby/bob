package schema

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type IndexBuilder struct {
	table   string
	name    string
	columns []string
	unique  bool
}

func newIndexBuilder() *IndexBuilder {
	return &IndexBuilder{
		columns: []string{},
	}
}

func (b *IndexBuilder) AddColumn(c string) *IndexBuilder {
	b.columns = append(b.columns, c)
	return b
}
func (b *IndexBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result().AddString("CREATE")
	if b.unique {
		r.AddString("UNIQUE")
	}
	r.AddString("INDEX IF NOT EXIST").
		Add(builder.Identifier(b.name)).
		AddString("ON").
		Add(builder.Identifier(b.table)).
		Add(builder.Group(builder.Join(builder.IdentifierList(b.columns), ", ")))

	return r.ToSQL(d)
}
