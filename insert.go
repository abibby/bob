package bob

import (
	"context"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
	"github.com/jmoiron/sqlx"
)

func SaveContext(ctx context.Context, tx *sqlx.Tx, v models.Model) error {
	return insert.SaveContext(ctx, tx, v)
}
func Save(tx *sqlx.Tx, v models.Model) error {
	return insert.Save(tx, v)
}
