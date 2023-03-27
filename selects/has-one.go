package selects

import (
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type HasOne[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[T]
}

var _ Relationship = &HasOne[models.Model]{}

func (r *HasOne[T]) Value(tx *sqlx.Tx) (T, error) {
	return r.relationValue.Value(tx, r)
}

//	func (r *HasOne[T]) Query() *selects.Builder {
//		var related T
//		local, ok := builder.GetValue(r.parent, r.localKey)
//		if !ok {
//			panic(fmt.Errorf("no local key %s", r.localKey))
//		}
//		return selects.New().
//			Select("*").
//			From(builder.GetTable(related)).
//			Where(r.foreignKey, "=", local).
//			Limit(1)
//	}
func (r *HasOne[T]) Loaded() bool {
	return r.loaded
}
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
func (r *HasOne[T]) Load(tx *sqlx.Tx, relations []Relationship) error {
	relatedLists, err := getRelated[T](tx, r, relations)
	if err != nil {
		return err
	}
	if len(relatedLists) == 0 {
		return nil
	}

	// TODO: replace with something more efficient
	for _, relation := range ofType[*HasOne[T]](relations) {
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
