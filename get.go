package bob

import (
	"github.com/abibby/bob/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

func (b *SelectBuilder) Get(tx *sqlx.Tx, v any) error {
	q, args, err := b.ToSQL(&mysql.MySQL{})
	if err != nil {
		return err
	}
	return tx.Select(v, q, args...)
}

func (b *SelectBuilder) First(tx *sqlx.Tx, v any) error {
	q, args, err := b.ToSQL(&mysql.MySQL{})
	if err != nil {
		return err
	}
	return tx.Get(v, q, args...)
}
