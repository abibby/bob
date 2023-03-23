package relationships

import (
	"reflect"

	"github.com/jmoiron/sqlx"
)

type HasMany[T any] struct {
	hasOneOrMany[T]
	value  []T
	loaded bool
}

var _ Relationship = &HasMany[any]{}

func (r *HasMany[T]) Value(tx *sqlx.Tx) ([]T, error) {
	if !r.loaded {
		err := r.Load(tx, []Relationship{r})
		if err != nil {
			return nil, err
		}
	}
	return r.value, nil
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
			foreign, foreignOk := getValue(related, r.getRelatedKey())
			if localOk && foreignOk && local == foreign {
				relation.value = append(relation.value, related)
			}

		}
		relation.loaded = true
	}
	return nil
}
