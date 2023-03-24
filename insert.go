package bob

import (
	"context"

	"github.com/abibby/bob/insert"
	"github.com/jmoiron/sqlx"
)

func SaveContext(ctx context.Context, tx *sqlx.Tx, v any) error {
	return insert.SaveContext(ctx, tx, v)
}
func Save(tx *sqlx.Tx, v any) error {
	return insert.Save(tx, v)
}
