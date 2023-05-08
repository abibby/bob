package bob

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
)

func SaveContext(ctx context.Context, tx builder.QueryExecer, v models.Model) error {
	return insert.SaveContext(ctx, tx, v)
}
func Save(tx builder.QueryExecer, v models.Model) error {
	return insert.Save(tx, v)
}
