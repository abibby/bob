package relationships

import (
	"reflect"

	"github.com/jmoiron/sqlx"
)

type HasOne[T any] struct {
	hasOneOrMany[T]
	value  T
	loaded bool
}

var _ Relationship = &HasOne[any]{}

func (r *HasOne[T]) Value(tx *sqlx.Tx) (T, error) {
	if !r.loaded {
		err := r.Load(tx, []Relationship{r})
		if err != nil {
			var zero T
			return zero, err
		}
	}
	return r.value, nil
}

//	func (r *HasOne[T]) Query() *selects.Builder {
//		var related T
//		local, ok := getValue(r.parent, r.localKey)
//		if !ok {
//			panic(fmt.Errorf("no local key %s", r.localKey))
//		}
//		return selects.New().
//			Select("*").
//			From(builder.GetTable(related)).
//			Where(r.foreignKey, "=", local).
//			Limit(1)
//	}
func (r *HasOne[T]) Initialize(parent any, field reflect.StructField) error {
	r.parent = parent
	r.parentKey = primaryKeyName(field, "local")
	r.relatedKey = foreignKeyName(field, "foreign", parent)
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
			foreign, foreignOk := getValue(related, r.getRelatedKey())
			if localOk && foreignOk && local == foreign {
				relation.value = related
			}

		}
		relation.loaded = true
	}

	return nil
}
