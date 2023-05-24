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

func InsertManyContext[T models.Model](ctx context.Context, tx builder.QueryExecer, models []T) error {
	return insert.InsertManyContext(ctx, tx, models)
}

func InsertMany[T models.Model](tx builder.QueryExecer, models []T) error {
	return insert.InsertMany(ctx, tx, models)
}
