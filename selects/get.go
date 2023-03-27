package selects

import (
	"context"
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/jmoiron/sqlx"
)

func (b *Builder[T]) Get(tx *sqlx.Tx) ([]T, error) {
	return b.GetContext(context.Background(), tx)
}
func (b *Builder[T]) GetContext(ctx context.Context, tx *sqlx.Tx) ([]T, error) {
	v := []T{}
	err := b.LoadContext(ctx, tx, &v)
	return v, err
}

func (b *Builder[T]) First(tx *sqlx.Tx) (T, error) {
	return b.FirstContext(context.Background(), tx)
}
func (b *Builder[T]) FirstContext(ctx context.Context, tx *sqlx.Tx) (T, error) {
	var v T
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		v = reflect.New(t.Elem()).Interface().(T)
	}
	err := b.LoadOneContext(ctx, tx, v)
	return v, err
}

func (b *Builder[T]) Find(tx *sqlx.Tx, primaryKeyValues ...any) (T, error) {
	return b.FindContext(context.Background(), tx)
}

func (b *Builder[T]) FindContext(ctx context.Context, tx *sqlx.Tx, primaryKeyValues ...any) (T, error) {
	var m T
	pKeys := builder.PrimaryKey(m)
	if len(pKeys) != len(primaryKeyValues) {
		return m, fmt.Errorf("")
	}
	for i, pKey := range pKeys {
		b.Where(pKey, "=", primaryKeyValues[i])
	}
	return b.FirstContext(ctx, tx)
}

func (b *Builder[T]) Load(tx *sqlx.Tx, v any) error {
	return b.LoadContext(context.Background(), tx, v)
}
func (b *Builder[T]) LoadContext(ctx context.Context, tx *sqlx.Tx, v any) error {
	q, bindings, err := b.ToSQL(dialects.DefaultDialect)
	if err != nil {
		return err
	}

	err = tx.SelectContext(ctx, v, q, bindings...)
	if err != nil {
		return err
	}

	err = InitializeRelationships(v)
	if err != nil {
		return err
	}

	err = hooks.AfterLoad(ctx, tx, v)
	if err != nil {
		return err
	}
	return nil
}

func (b *Builder[T]) LoadOne(tx *sqlx.Tx, v any) error {
	return b.LoadOneContext(context.Background(), tx, v)
}
func (b *Builder[T]) LoadOneContext(ctx context.Context, tx *sqlx.Tx, v any) error {
	q, bindings, err := b.Clone().Limit(1).ToSQL(dialects.DefaultDialect)

	if err != nil {
		return err
	}

	err = tx.GetContext(ctx, v, q, bindings...)
	if err != nil {
		return err
	}
	err = InitializeRelationships(v)
	if err != nil {
		return err
	}

	err = hooks.AfterLoad(ctx, tx, v)
	if err != nil {
		return err
	}
	return nil
}
