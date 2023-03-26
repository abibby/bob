package selects

import (
	"context"
	"reflect"

	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/jmoiron/sqlx"
)

func (b *Builder[T]) Get(tx *sqlx.Tx) ([]T, error) {
	return b.GetContext(context.Background(), tx)
}
func (b *Builder[T]) GetContext(ctx context.Context, tx *sqlx.Tx) ([]T, error) {
	q, bindings, err := b.ToSQL(dialects.DefaultDialect)
	if err != nil {
		return nil, err
	}

	v := []T{}
	err = tx.SelectContext(ctx, &v, q, bindings...)
	if err != nil {
		return nil, err
	}

	err = InitializeRelationships(&v)
	if err != nil {
		return nil, err
	}

	err = hooks.AfterLoad(ctx, tx, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (b *Builder[T]) First(tx *sqlx.Tx) (T, error) {
	return b.FirstContext(context.Background(), tx)
}
func (b *Builder[T]) FirstContext(ctx context.Context, tx *sqlx.Tx) (T, error) {
	lastLimit := b.limit
	q, bindings, err := b.Limit(1).ToSQL(dialects.DefaultDialect)
	b.limit = lastLimit

	if err != nil {
		var zero T
		return zero, err
	}

	var v T
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		v = reflect.New(t.Elem()).Interface().(T)
	}
	err = tx.GetContext(ctx, v, q, bindings...)
	if err != nil {
		var zero T
		return zero, err
	}
	err = InitializeRelationships(v)
	if err != nil {
		var zero T
		return zero, err
	}

	err = hooks.AfterLoad(ctx, tx, v)
	if err != nil {
		var zero T
		return zero, err
	}
	return v, nil
}
