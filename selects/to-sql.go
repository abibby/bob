package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

func (b *Builder[T]) ToSQL(d dialects.Dialect) (string, []any, error) {
	var m T
	var s any = m
	b = b.Clone()
	if scoper, ok := s.(Scoper); ok {
		for _, scope := range scoper.Scopes() {
			if !b.withoutScopes.Has(scope.Name) {
				b.subBuilder = scope.Apply(b.subBuilder)
			}
		}
	}
	return b.subBuilder.ToSQL(d)
}
func (b *SubBuilder) ToSQL(d dialects.Dialect) (string, []any, error) {
	b = b.Clone()
	for _, scope := range b.scopes {
		b = scope.Apply(b)
	}
	return builder.Result().
		Add(b.selects.ToSQL(d)).
		Add(b.from.ToSQL(d)).
		Add(b.wheres.ToSQL(d)).
		Add(b.groupBys.ToSQL(d)).
		Add(b.havings.ToSQL(d)).
		Add(b.orderBys.ToSQL(d)).
		Add(b.limit.ToSQL(d)).
		ToSQL(d)
}
