package insert

import (
	"context"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/abibby/bob/selects"
	scan "github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

func Save(tx *sqlx.Tx, v any) error {
	return SaveContext(context.Background(), tx, v)
}
func SaveContext(ctx context.Context, tx *sqlx.Tx, v any) error {
	err := hooks.BeforeSave(ctx, tx, v)
	if err != nil {
		return err
	}

	d := dialects.DefaultDialect
	columns, err := scan.Columns(v)
	if err != nil {
		return err
	}
	values, err := scan.Values(columns, v)
	if err != nil {
		return err
	}

	err = insert(ctx, tx, d, v, columns, values)
	if err != nil {
		return err
	}

	err = selects.InitializeRelationships(v)
	if err != nil {
		return err
	}
	return hooks.AfterSave(ctx, tx, v)
}

func insert(ctx context.Context, tx *sqlx.Tx, d dialects.Dialect, v any, columns []string, values []any) error {

	r := builder.Result().
		AddString("INSERT INTO").
		Add(builder.Identifier(builder.GetTable(v)).ToSQL(d)).
		Add(
			builder.NewGroup(
				builder.Join(
					builder.IdentifierList(columns),
					", ",
				),
			).ToSQL(d),
		).
		AddString("VALUES").
		Add(
			builder.NewGroup(
				builder.Join(
					builder.LiteralList(values),
					", ",
				),
			).ToSQL(d),
		)

	q, bindings, err := r.ToSQL(d)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, q, bindings...)
	if err != nil {
		return err
	}
	return nil
}
