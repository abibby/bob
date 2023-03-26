package selects

import (
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type BelongsTo[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[T]
}

var _ Relationship = &BelongsTo[models.Model]{}

func (r *BelongsTo[T]) Value(tx *sqlx.Tx) (T, error) {
	return r.relationValue.Value(tx, r)
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
			foreign, foreignOk := builder.GetValue(related, r.getRelatedKey())
			if localOk && foreignOk && local == foreign {
				relation.value = related
			}
		}
		relation.loaded = true
	}

	return nil
}
