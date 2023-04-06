package selects

import (
	"context"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type HasMany[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[[]T]
}

var _ Relationship = &HasMany[models.Model]{}

func (r *HasMany[T]) Initialize(parent any, field reflect.StructField) error {
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
func (r *HasMany[T]) Load(ctx context.Context, tx *sqlx.Tx, relations []Relationship) error {
	relatedLists, err := r.getRelated(ctx, tx, relations)
	if err != nil {
		return err
	}
	if len(relatedLists) == 0 {
		return nil
	}

	// TODO: replace with something more efficient
	for _, relation := range ofType[*HasMany[T]](relations) {
		relation.value = []T{}
		for _, related := range relatedLists {
			local, localOk := relation.parentKeyValue()
			foreign, foreignOk := builder.GetValue(related, r.getRelatedKey())
			if localOk && foreignOk && local == foreign {
				relation.value = append(relation.value, related)
			}

		}
		relation.loaded = true
	}
	return nil
}
