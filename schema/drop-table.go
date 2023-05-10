package schema

import (
	"context"

	"github.com/abibby/bob/builder"
)

func Drop(table string) Runner {
	return Run(func(ctx context.Context, tx builder.QueryExecer) error {
		return runQuery(ctx, tx, builder.Concat(builder.Raw("DROP TABLE "), builder.Identifier(table)))
	})
}
func DropIfExists(table string) Runner {
	return Run(func(ctx context.Context, tx builder.QueryExecer) error {
		return runQuery(ctx, tx, builder.Concat(builder.Raw("DROP TABLE IF EXISTS "), builder.Identifier(table)))
	})
}
