package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

type iHasOneOrMany interface {
	getParentKey() string
	getRelatedKey() string
	parentKeyValue() (any, bool)
	relatedKeyValue() (any, bool)
}

type hasOneOrMany[T models.Model] struct {
	parent     any
	relatedKey string
	parentKey  string
}

var _ iHasOneOrMany = hasOneOrMany[models.Model]{}

func (r hasOneOrMany[T]) parentKeyValue() (any, bool) {
	return builder.GetValue(r.parent, r.parentKey)
}
func (r hasOneOrMany[T]) relatedKeyValue() (any, bool) {
	var related T
	return builder.GetValue(related, r.relatedKey)
}

func (r hasOneOrMany[T]) getParentKey() string {
	return r.parentKey
}
func (r hasOneOrMany[T]) getRelatedKey() string {
	return r.relatedKey
}

func getRelated[T models.Model](tx *sqlx.Tx, r iHasOneOrMany, relations []Relationship) ([]T, error) {
	var related T
	localKeys := make([]any, 0, len(relations))
	for _, r := range relations {
		local, ok := r.(iHasOneOrMany).parentKeyValue()
		if !ok {
			continue
		}
		localKeys = append(localKeys, local)
	}

	return New[T]().
		Select("*").
		From(builder.GetTable(related)).
		WhereIn(r.getRelatedKey(), localKeys).
		Get(tx)
}
