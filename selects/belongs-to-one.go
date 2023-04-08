package selects

import (
	"context"
	"reflect"

	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type BelongsTo[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[T]
}

var _ Relationship = &BelongsTo[models.Model]{}

func (r *BelongsTo[T]) Initialize(parent any, field reflect.StructField) error {
	var related T
	r.parent = parent
	parentKey, err := foreignKeyName(field, "foreign", related)
	if err != nil {
		return err
	}
	relatedKey, err := primaryKeyName(field, "owner", related)
	if err != nil {
		return err
	}

	r.parentKey = parentKey
	r.relatedKey = relatedKey

	return nil
}

// func (r *BelongsTo[T]) Load(ctx context.Context, tx *sqlx.Tx, relations []Relationship) error {
// 	relatedLists, err := r.getRelated(ctx, tx, relations)
// 	if err != nil {
// 		return err
// 	}
// 	if len(relatedLists) == 0 {
// 		return nil
// 	}

// 	// TODO: replace with something more efficient
// 	for _, relation := range ofType[*BelongsTo[T]](relations) {
// 		for _, related := range relatedLists {
// 			local, localOk := relation.parentKeyValue()
// 			foreign, foreignOk := builder.GetValue(related, r.getRelatedKey())
// 			if localOk && foreignOk && local == foreign {
// 				relation.value = related
// 			}
// 		}
// 		relation.loaded = true
// 	}

// 	return nil
// }

func (r *BelongsTo[T]) Load(ctx context.Context, tx *sqlx.Tx, relations []Relationship) error {
	relatedMap, err := r.relatedMap(ctx, tx, relations)
	if err != nil {
		return err
	}

	for _, relation := range ofType[*BelongsTo[T]](relations) {
		local, ok := relation.parentKeyValue()
		if !ok {
			continue
		}
		m, ok := relatedMap[local]
		if ok {
			relation.value = m[0]
		}
		relation.loaded = true
	}
	return nil
}
