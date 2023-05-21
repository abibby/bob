package schema

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type Runner interface {
	Run(ctx context.Context, tx builder.QueryExecer) error
}

type RunnerFunc func(ctx context.Context, tx builder.QueryExecer) error

func (f RunnerFunc) Run(ctx context.Context, tx builder.QueryExecer) error {
	return f(ctx, tx)
}

func Run(f RunnerFunc) Runner {
	return f
}

func runQuery(ctx context.Context, tx builder.QueryExecer, sqler builder.ToSQLer) error {
	sql, bindings, err := sqler.ToSQL(dialects.New())
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, sql, bindings...)
	return err
}
