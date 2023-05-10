package selects

import (
	"context"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

// BelongsTo represents a belongs to relationship on a model. The parent model
// with a BelongsTo property will have a column referencing another tables
// primary key. For example if model Foo had a BelongsTo[*Bar] property the foos
// table would have a foos.bar_id column related to the bars.id column. Struct
// tags can be used to change the column names if they don't follow the default
// naming convention. The column on the parent model can be set with a foreign
// tag and the column on the related model can be set with an owner tag.
type BelongsTo[T models.Model] struct {
	hasOneOrMany[T]
	relationValue[T]
}

var _ Relationship = &BelongsTo[models.Model]{}

func (r *BelongsTo[T]) Initialize(parent any, field reflect.StructField) error {
	var related T
	r.parent = parent
	parentKey, err := foreignKeyName(field, "foreign", related)
	if err != nil {
		return err
	}
	relatedKey, err := primaryKeyName(field, "owner", related)
	if err != nil {
		return err
	}

	r.parentKey = parentKey
	r.relatedKey = relatedKey

	return nil
}

func (r *BelongsTo[T]) Load(ctx context.Context, tx builder.QueryExecer, relations []Relationship) error {
	rm, err := r.relatedMap(ctx, tx, relations)
	if err != nil {
		return err
	}

	for _, relation := range ofType[*BelongsTo[T]](relations) {
		relation.value = rm.Single(relation.parentKeyValue())
		relation.loaded = true
	}
	return nil
}
