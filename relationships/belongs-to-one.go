package relationships

import (
	"reflect"

	"github.com/jmoiron/sqlx"
)

type BelongsTo[T any] struct {
	hasOneOrMany[T]
	value  T
	loaded bool
}

var _ Relationship = &BelongsTo[any]{}

func (r *BelongsTo[T]) Value(tx *sqlx.Tx) (T, error) {
	if !r.loaded {
		err := r.Load(tx, []Relationship{r})
		if err != nil {
			var zero T
			return zero, err
		}
	}
	return r.value, nil
}

func (r *BelongsTo[T]) Initialize(parent any, field reflect.StructField) error {
	var related T
	r.parent = parent
	r.parentKey = foreignKeyName(field, "foreign", related)
	r.relatedKey = primaryKeyName(field, "owner")
	return nil
}

func (r *BelongsTo[T]) Load(tx *sqlx.Tx, relations []Relationship) error {
	relatedLists, err := getRelated[T](tx, r, relations)
	if err != nil {
		return err
	}
	if len(relatedLists) == 0 {
		return nil
	}

	// TODO: replace with something more efficient
	for _, relation := range ofType[*BelongsTo[T]](relations) {
		for _, related := range relatedLists {
			local, localOk := relation.parentKeyValue()
			foreign, foreignOk := getValue(related, r.getRelatedKey())
			if localOk && foreignOk && local == foreign {
				relation.value = related
			}
		}
		relation.loaded = true
	}

	return nil
}
