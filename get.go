package bob

import (
	"github.com/abibby/bob/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

func (b *SelectBuilder) Get(v any) error {
	q, args, err := b.ToSQL(&mysql.MySQL{})
	if err != nil {
		return err
	}
	return sqlx.Select(nil, v, q, args...)
}

func (b *SelectBuilder) First(v any) error {
	q, args, err := b.ToSQL(&mysql.MySQL{})
	if err != nil {
		return err
	}
	return sqlx.Get(nil, v, q, args...)
}
