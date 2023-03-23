package relationships

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects/mysql"
	"github.com/abibby/bob/selects"
	"github.com/jmoiron/sqlx"
)

type iHasOneOrMany interface {
	getParentKey() string
	getRelatedKey() string
	parentKeyValue() (any, bool)
	relatedKeyValue() (any, bool)
}

type hasOneOrMany[T any] struct {
	parent     any
	relatedKey string
	parentKey  string
}

var _ iHasOneOrMany = hasOneOrMany[any]{}

func (r hasOneOrMany[T]) parentKeyValue() (any, bool) {
	return getValue(r.parent, r.parentKey)
}
func (r hasOneOrMany[T]) relatedKeyValue() (any, bool) {
	var related T
	return getValue(related, r.relatedKey)
}

func (r hasOneOrMany[T]) getParentKey() string {
	return r.parentKey
}
func (r hasOneOrMany[T]) getRelatedKey() string {
	return r.relatedKey
}

func getRelated[T any](tx *sqlx.Tx, r iHasOneOrMany, relations []Relationship) ([]T, error) {
	var related T
	localKeys := make([]any, 0, len(relations))
	for _, r := range relations {
		local, ok := r.(iHasOneOrMany).parentKeyValue()
		if !ok {
			continue
		}
		localKeys = append(localKeys, local)
	}

	relatedLists := []T{}

	err := selects.New().
		Select("*").
		From(builder.GetTable(related)).
		WhereIn(r.getRelatedKey(), localKeys).
		Dump(&mysql.MySQL{}).
		Get(tx, &relatedLists)
	if err != nil {
		return nil, err
	}
	return relatedLists, nil
}
