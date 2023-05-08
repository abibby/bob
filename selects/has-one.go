package selects

import (
	"context"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

// HasOne
// tags `local` `foreign`
type HasOne[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[T]
}

var _ Relationship = &HasOne[models.Model]{}

func (r *HasOne[T]) Initialize(parent any, field reflect.StructField) error {
	r.parent = parent
	parentKey, err := primaryKeyName(field, "local", parent)
	if err != nil {
		return err
	}
	relatedKey, err := foreignKeyName(field, "foreign", parent)
	if err != nil {
		return err
	}
	r.parentKey = parentKey
	r.relatedKey = relatedKey
	return nil
}

// func (r *HasOne[T]) Load(ctx context.Context, tx builder.QueryExecer, relations []Relationship) error {
// 	relatedLists, err := r.getRelated(ctx, tx, relations)
// 	if err != nil {
// 		return err
// 	}
// 	if len(relatedLists) == 0 {
// 		return nil
// 	}

// 	// TODO: replace with something more efficient
// 	for _, relation := range ofType[*HasOne[T]](relations) {
// 		local, ok := relation.parentKeyValue()
// 		if !ok {
// 			continue
// 		}
// 		for _, related := range relatedLists {
// 			foreign, ok := builder.GetValue(related, r.getRelatedKey())
// 			if !ok {
// 				continue
// 			}
// 			if local == foreign {
// 				relation.value = related
// 			}

// 		}
// 		relation.loaded = true
// 	}

//		return nil
//	}
func (r *HasOne[T]) Load(ctx context.Context, tx builder.QueryExecer, relations []Relationship) error {
	rm, err := r.relatedMap(ctx, tx, relations)
	if err != nil {
		return err
	}

	for _, relation := range ofType[*HasOne[T]](relations) {
		relation.value = rm.Single(relation.parentKeyValue())
		relation.loaded = true
	}
	return nil
}
