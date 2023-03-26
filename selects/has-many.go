package selects

import (
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

func (r *HasMany[T]) Value(tx *sqlx.Tx) ([]T, error) {
	return r.relationValue.Value(tx, r)
}

func (r *HasMany[T]) Loaded() bool {
	return r.loaded
}
func (r *HasMany[T]) Initialize(parent any, field reflect.StructField) error {
	r.parent = parent
	r.parentKey = primaryKeyName(field, "local")
	r.relatedKey = foreignKeyName(field, "foreign", parent)
	return nil
}
func (r *HasMany[T]) Load(tx *sqlx.Tx, relations []Relationship) error {
	relatedLists, err := getRelated[T](tx, r, relations)
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
