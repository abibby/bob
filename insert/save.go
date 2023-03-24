package insert

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/selects"
	scan "github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

func Save(tx *sqlx.Tx, v any) error {
	d := dialects.DefaultDialect
	columns, err := scan.Columns(v)
	if err != nil {
		return err
	}
	values, err := scan.Values(columns, v)
	if err != nil {
		return err
	}
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
	_, err = tx.Exec(q, bindings...)
	if err != nil {
		return err
	}
	return selects.InitializeRelationships(v)
}
