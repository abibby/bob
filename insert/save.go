package insert

import (
	"context"
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/hooks"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
	"github.com/jmoiron/sqlx"
)

func columnsAndValues(v reflect.Value) ([]string, []any) {
	t := v.Type()
	numFields := t.NumField()
	columns := make([]string, 0, numFields)
	values := make([]any, 0, numFields)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		if field.Anonymous {
			subColumns, subValues := columnsAndValues(v.Field(i))
			columns = append(columns, subColumns...)
			values = append(values, subValues...)
		} else {
			name := builder.FieldName(field)
			if name == "-" {
				continue
			}
			columns = append(columns, name)
			values = append(values, v.Field(i).Interface())
		}
	}
	return columns, values
}

func Save(tx *sqlx.Tx, v models.Model) error {
	return SaveContext(context.Background(), tx, v)
}
func SaveContext(ctx context.Context, tx *sqlx.Tx, v models.Model) error {
	err := hooks.BeforeSave(ctx, tx, v)
	if err != nil {
		return err
	}

	d := dialects.DefaultDialect
	columns, values := columnsAndValues(reflect.ValueOf(v).Elem())

	if v.InDatabase() {
		err = update(ctx, tx, d, v, columns, values)
		if err != nil {
			return err
		}
	} else {
		err = insert(ctx, tx, d, v, columns, values)
		if err != nil {
			return err
		}
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
		Add(builder.Identifier(builder.GetTable(v)).ToSQL(ctx, d)).
		Add(
			builder.Group(
				builder.Join(
					builder.IdentifierList(columns),
					", ",
				),
			).ToSQL(ctx, d),
		).
		AddString("VALUES").
		Add(
			builder.Group(
				builder.Join(
					builder.LiteralList(values),
					", ",
				),
			).ToSQL(ctx, d),
		)

	q, bindings, err := r.ToSQL(ctx, d)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, q, bindings...)
	if err != nil {
		return err
	}
	return nil
}

func update(ctx context.Context, tx *sqlx.Tx, d dialects.Dialect, v any, columns []string, values []any) error {
	pKey := builder.PrimaryKey(v)
	r := builder.Result().
		AddString("UPDATE").
		Add(builder.Identifier(builder.GetTable(v)).ToSQL(ctx, d)).
		AddString("SET")

	for i, column := range columns {
		if i != 0 {
			r.AddString(",")
		}
		r.Add(builder.Identifier(column).ToSQL(ctx, d))
		r.AddString("=")
		r.Add(builder.Literal(values[i]).ToSQL(ctx, d))
	}

	r.AddString("WHERE")

	for i, k := range pKey {
		pKeyValue, ok := builder.GetValue(v, k)
		if !ok {
			return fmt.Errorf("no primary key found")
		}

		if i != 0 {
			r.AddString("AND")
		}

		r.Add(builder.Identifier(k).ToSQL(ctx, d)).
			AddString("=").
			Add(builder.Literal(pKeyValue).ToSQL(ctx, d))
	}

	q, bindings, err := r.ToSQL(ctx, d)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, q, bindings...)
	if err != nil {
		return err
	}
	return nil
}
