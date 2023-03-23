package selects

import (
	"github.com/abibby/bob/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

func (b *Builder) Get(tx *sqlx.Tx, v any) error {
	q, bindings, err := b.ToSQL(&mysql.MySQL{})
	if err != nil {
		return err
	}

	return tx.Select(v, q, bindings...)
}

func (b *Builder) First(tx *sqlx.Tx, v any) error {
	lastLimit := b.limit
	q, bindings, err := b.Limit(1).ToSQL(&mysql.MySQL{})
	b.limit = lastLimit

	if err != nil {
		return err
	}

	return tx.Get(v, q, bindings...)
}
