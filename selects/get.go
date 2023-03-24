package selects

import (
	"context"

	"github.com/abibby/bob/dialects"
	"github.com/jmoiron/sqlx"
)

func (b *Builder) Get(tx *sqlx.Tx, v any) error {
	return b.GetContext(context.Background(), tx, v)
}
func (b *Builder) GetContext(ctx context.Context, tx *sqlx.Tx, v any) error {
	q, bindings, err := b.ToSQL(dialects.DefaultDialect)
	if err != nil {
		return err
	}

	err = tx.SelectContext(ctx, v, q, bindings...)
	if err != nil {
		return err
	}
	return InitializeRelationships(v)
}

func (b *Builder) First(tx *sqlx.Tx, v any) error {
	return b.FirstContext(context.Background(), tx, v)
}
func (b *Builder) FirstContext(ctx context.Context, tx *sqlx.Tx, v any) error {
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
	return InitializeRelationships(v)
}
