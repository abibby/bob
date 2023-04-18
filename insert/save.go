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

var relationshipInterface = reflect.TypeOf((*selects.Relationship)(nil)).Elem()

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
			tag := builder.DBTag(field)
			if tag.Name == "-" || tag.Readonly || field.Type.Implements(relationshipInterface) {
				continue
			}
			columns = append(columns, tag.Name)
			values = append(values, v.Field(i).Interface())
		}
	}
	return columns, values
}

func Save(tx *sqlx.Tx, v models.Model) error {
	ctx := context.Background()
	if v, ok := v.(models.Contexter); ok {
		modelCtx := v.Context()
		if modelCtx != nil {
			ctx = modelCtx
		}
	}
	return SaveContext(ctx, tx, v)
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
	rPKey, pKey, isAuto := isAutoIncrementing(v)
	if isAuto {
		newColumns := make([]string, 0, len(columns))
		newValues := make([]any, 0, len(values))
		for i, column := range columns {
			if column != pKey {
				newColumns = append(newColumns, column)
				newValues = append(newValues, values[i])
			}
		}
		columns = newColumns
		values = newValues
	}
	r := builder.Result().
		AddString("INSERT INTO").
		Add(builder.Identifier(builder.GetTable(v))).
		Add(
			builder.Group(
				builder.Join(
					builder.IdentifierList(columns),
					", ",
				),
			),
		).
		AddString("VALUES").
		Add(
			builder.Group(
				builder.Join(
					builder.LiteralList(values),
					", ",
				),
			),
		)

	q, bindings, err := r.ToSQL(d)
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, q, bindings...)
	if err != nil {
		return err
	}

	if isAuto {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		rPKey.SetInt(id)
	}
	return nil
}

func isAutoIncrementing(v any) (reflect.Value, string, bool) {
	pKeys := builder.PrimaryKey(v)
	if len(pKeys) != 1 {
		return reflect.Value{}, "", false
	}

	pKey := pKeys[0]
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return reflect.Value{}, "", false
	}
	var pKeyTag *builder.Tag
	var rPKey reflect.Value
	errFound := fmt.Errorf("found")
	err := builder.EachField(rv, func(sf reflect.StructField, fv reflect.Value) error {
		tag := builder.DBTag(sf)
		if tag.Name == pKey {
			pKeyTag = tag
			rPKey = fv
			if !rPKey.IsZero() {
				return nil
			}
			return errFound
		}
		return nil
	})
	if err != errFound {
		return reflect.Value{}, "", false
	}
	if pKeyTag != nil && !pKeyTag.AutoIncrement {
		return reflect.Value{}, "", false
	}
	return rPKey, pKey, true
}

func update(ctx context.Context, tx *sqlx.Tx, d dialects.Dialect, v any, columns []string, values []any) error {
	pKey := builder.PrimaryKey(v)
	r := builder.Result().
		AddString("UPDATE").
		Add(builder.Identifier(builder.GetTable(v))).
		AddString("SET")

	for i, column := range columns {
		if i != 0 {
			r.AddString(",")
		}
		r.Add(builder.Identifier(column))
		r.AddString("=")
		r.Add(builder.Literal(values[i]))
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

		r.Add(builder.Identifier(k)).
			AddString("=").
			Add(builder.Literal(pKeyValue))
	}

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
