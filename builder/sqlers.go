package builder

import (
	"context"

	"github.com/abibby/bob/dialects"
)

type ToSQLer interface {
	ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error)
}

type ToSQLFunc func(ctx context.Context, d dialects.Dialect) (string, []any, error)

func (f ToSQLFunc) ToSQL(ctx context.Context, d dialects.Dialect) (string, []any, error) {
	return f(ctx, d)
}

func Identifier(i string) ToSQLer {
	return ToSQLFunc(func(ctx context.Context, d dialects.Dialect) (string, []any, error) {
		return d.Identifier(i), nil, nil
	})
}

func IdentifierList(strs []string) []ToSQLer {
	identifiers := make([]ToSQLer, len(strs))
	for i, s := range strs {
		identifiers[i] = Identifier(s)
	}
	return identifiers
}

func Join(sqlers []ToSQLer, sep string) ToSQLer {
	return ToSQLFunc(func(ctx context.Context, d dialects.Dialect) (string, []any, error) {
		sql := ""
		bindings := []any{}
		for i, sqler := range sqlers {
			q, b, err := sqler.ToSQL(ctx, d)
			if err != nil {
				return "", nil, err
			}
			sql += q
			if i < len(sqlers)-1 {
				sql += sep
			}
			bindings = append(bindings, b...)
		}
		return sql, bindings, nil
	})
}

func Raw(sql string, bindings ...any) ToSQLer {
	return ToSQLFunc(func(ctx context.Context, d dialects.Dialect) (string, []any, error) {
		return sql, bindings, nil
	})
}

func Group(sqler ToSQLer) ToSQLer {
	return ToSQLFunc(func(ctx context.Context, d dialects.Dialect) (string, []any, error) {
		q, bindings, err := sqler.ToSQL(ctx, d)
		return "(" + q + ")", bindings, err
	})
}

func Literal(v any) ToSQLer {
	return ToSQLFunc(func(ctx context.Context, d dialects.Dialect) (string, []any, error) {
		return "?", []any{v}, nil
	})
}

func LiteralList(values []any) []ToSQLer {
	literals := make([]ToSQLer, len(values))
	for i, s := range values {
		literals[i] = Literal(s)
	}
	return literals
}
