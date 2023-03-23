package relationships

import (
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/selects"
	"github.com/jmoiron/sqlx"
)

type HasOne[T any] struct {
	parent     any
	foreignKey string
	localKey   string
	value      T
	loaded     bool
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
	r.localKey, _ = field.Tag.Lookup("local")
	r.foreignKey, _ = field.Tag.Lookup("foreign")
	return nil
}
func (r *HasOne[T]) Load(tx *sqlx.Tx, relations []Relationship) error {
	var related T
	hasOnes := ofType[*HasOne[T]](relations)
	localKeys := make([]any, 0, len(hasOnes))
	for _, r := range hasOnes {
		local, ok := getValue(r.parent, r.localKey)
		if !ok {
			continue
		}
		localKeys = append(localKeys, local)
	}

	relatedLists := []T{}

	err := selects.New().
		Select("*").
		From(builder.GetTable(related)).
		WhereIn(r.foreignKey, localKeys).
		Get(tx, &relatedLists)
	if err != nil {
		return err
	}

	if len(relatedLists) == 0 {
		return nil
	}

	// TODO: replace with something more efficient
	for _, relation := range hasOnes {
		for _, related := range relatedLists {
			local, localOk := getValue(relation.parent, r.localKey)
			foreign, foreignOk := getValue(related, r.foreignKey)
			if localOk && foreignOk && local == foreign {
				relation.value = related
			}

		}
	}

	return nil
}
