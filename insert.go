package bob

import (
	"github.com/abibby/bob/insert"
	"github.com/jmoiron/sqlx"
)

func Save(tx *sqlx.Tx, v any) error {
	return insert.Save(tx, v)
}
