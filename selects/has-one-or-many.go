package selects

import (
	"context"
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type iHasOneOrMany interface {
	getParentKey() string
	getRelatedKey() string
	parentKeyValue() (any, bool)
	relatedKeyValue() (any, bool)
}

type hasOneOrMany[T models.Model] struct {
	parent     any
	relatedKey string
	parentKey  string
}

var _ iHasOneOrMany = hasOneOrMany[models.Model]{}

func (r hasOneOrMany[T]) Subquery() *SubBuilder {
	return From[T]().
		WhereColumn(r.relatedKey, "=", builder.GetTable(r.parent)+"."+r.parentKey).
		subBuilder
}

func (r hasOneOrMany[T]) parentKeyValue() (any, bool) {
	return builder.GetValue(r.parent, r.parentKey)
}
func (r hasOneOrMany[T]) relatedKeyValue() (any, bool) {
	var related T
	return builder.GetValue(related, r.relatedKey)
}

func (r hasOneOrMany[T]) getParentKey() string {
	return r.parentKey
}
func (r hasOneOrMany[T]) getRelatedKey() string {
	return r.relatedKey
}

func (r hasOneOrMany[T]) getRelated(ctx context.Context, tx *sqlx.Tx, relations []Relationship) ([]T, error) {
	localKeys := make([]any, 0, len(relations))
	for _, r := range relations {
		local, ok := r.(iHasOneOrMany).parentKeyValue()
		if !ok {
			var related T
			return nil, fmt.Errorf("%s has no field %s: %w", reflect.TypeOf(related).Name(), r.(iHasOneOrMany).getParentKey(), ErrMissingField)
		}
		if local != nil {
			localKeys = append(localKeys, local)
		}
	}

	return From[T]().
		WhereIn(r.getRelatedKey(), localKeys).
		WithContext(ctx).
		Get(tx)
}
func (r hasOneOrMany[T]) relatedMap(ctx context.Context, tx *sqlx.Tx, relations []Relationship) (map[any][]T, error) {
	var related T
	if !builder.HasField(related, r.getRelatedKey()) {
		return nil, fmt.Errorf("%s has no field %s: %w", reflect.TypeOf(related).Name(), r.getRelatedKey(), ErrMissingField)
	}

	relatedLists, err := r.getRelated(ctx, tx, relations)
	if err != nil {
		return nil, err
	}
	relatedMap := map[any][]T{}
	for _, related := range relatedLists {
		foreign, ok := builder.GetValue(related, r.getRelatedKey())
		if !ok {
			return nil, fmt.Errorf("%s has no field %s: %w", reflect.TypeOf(related).Name(), r.getRelatedKey(), ErrMissingField)
		}
		if str, ok := foreign.(fmt.Stringer); ok {
			foreign = str.String()
		}
		m, ok := relatedMap[foreign]
		if !ok {
			m = []T{related}
		} else {
			m = append(m, related)
		}
		relatedMap[foreign] = m
	}

	return relatedMap, nil
}
