package migrate

import (
	"context"
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

func RunModelCreate(ctx context.Context, db builder.QueryExecer, models ...models.Model) error {
	for _, model := range models {
		m, err := CreateFromModel(model)
		if err != nil {
			return fmt.Errorf("migration for %s: %w", builder.GetTable(model), err)
		}
		err = m.Run(ctx, db)
		if err != nil {
			return fmt.Errorf("migration for %s: %w", builder.GetTable(model), err)
		}
	}
	return nil
}
