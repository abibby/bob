package craetetable

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type ForeignKeyBuilder struct {
	relatedTable string
	localKey     string
	relatedKey   string
}

func (b *ForeignKeyBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := builder.Result()
	r.AddString("CONSTRAINT").
		Add(builder.Identifier(b.relatedTable)).
		AddString("FOREIGN KEY").
		Add(builder.Group(builder.Identifier(b.localKey))).
		AddString("REFERENCES").
		Add(builder.Concat(
			builder.Identifier(b.relatedTable),
			builder.Group(builder.Identifier(b.relatedKey)),
		))
	return r.ToSQL(d)
}
