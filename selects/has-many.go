package selects

import (
	"context"
	"reflect"

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
	rm, err := r.relatedMap(ctx, tx, relations)
	if err != nil {
		return err
	}

	for _, relation := range ofType[*HasMany[T]](relations) {
		relation.value = rm.Multi(relation.parentKeyValue())
		relation.loaded = true
	}
	return nil
}
