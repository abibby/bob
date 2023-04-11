package selects

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/jmoiron/sqlx"
)

func (b *Builder[T]) Get(tx *sqlx.Tx) ([]T, error) {
	v := []T{}
	err := b.Load(tx, &v)
	if err != nil {
		return nil, err
	}

	for _, with := range b.withs {
		err = LoadMissingContext(b.Context(), tx, v, with)
		if err != nil {
			return nil, err
		}
	}

	return v, nil
}

func (b *Builder[T]) First(tx *sqlx.Tx) (T, error) {
	v, err := b.Clone().
		Limit(1).
		Get(tx)
	if err != nil {
		var zero T
		return zero, err
	}
	if len(v) < 1 {
		var zero T
		return zero, nil
	}
	return v[0], err
}

func (b *Builder[T]) Find(tx *sqlx.Tx, primaryKeyValue any) (T, error) {
	var m T
	pKeys := builder.PrimaryKey(m)
	if len(pKeys) != 1 {
		return m, fmt.Errorf("Find only supports tables with 1 primary key")
	}
	return b.Clone().
		Where(pKeys[0], "=", primaryKeyValue).
		First(tx)
}

func (b *Builder[T]) Load(tx *sqlx.Tx, v any) error {
	q, bindings, err := b.ToSQL(dialects.DefaultDialect)
	if err != nil {
		return err
	}

	err = tx.SelectContext(b.Context(), v, q, bindings...)
	if err != nil {
		return err
	}

	err = InitializeRelationships(v)
	if err != nil {
		return err
	}

	err = hooks.AfterLoad(b.Context(), tx, v)
	if err != nil {
		return err
	}
	return nil
}

func (b *Builder[T]) LoadOne(tx *sqlx.Tx, v any) error {
	q, bindings, err := b.Clone().
		Limit(1).
		ToSQL(dialects.DefaultDialect)

	if err != nil {
		return err
	}

	err = tx.GetContext(b.Context(), v, q, bindings...)
	if err != nil {
		return err
	}
	err = InitializeRelationships(v)
	if err != nil {
		return err
	}

	err = hooks.AfterLoad(b.Context(), tx, v)
	if err != nil {
		return err
	}
	return nil
}
