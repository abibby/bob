package selects

import (
	"context"

	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/jmoiron/sqlx"
)

func (b *Builder[T]) Get(tx *sqlx.Tx, v *[]T) error {
	return b.GetContext(context.Background(), tx, v)
}
func (b *Builder[T]) GetContext(ctx context.Context, tx *sqlx.Tx, v *[]T) error {
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

	return hooks.AfterLoad(ctx, tx, v)
}

func (b *Builder[T]) First(tx *sqlx.Tx, v T) error {
	return b.FirstContext(context.Background(), tx, v)
}
func (b *Builder[T]) FirstContext(ctx context.Context, tx *sqlx.Tx, v T) error {
	lastLimit := b.limit
	q, bindings, err := b.Limit(1).ToSQL(dialects.DefaultDialect)
	b.limit = lastLimit

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

	return hooks.AfterLoad(ctx, tx, v)
}
