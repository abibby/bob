package builder

import "github.com/abibby/bob/dialects"

type SQLResult struct {
	sql      string
	bindings []any
	err      error
}

func Result() *SQLResult {
	return &SQLResult{}
}

func (r *SQLResult) AddString(sql string) *SQLResult {
	return r.Add(sql, nil, nil)
}
func (r *SQLResult) Add(sql string, bindings []any, err error) *SQLResult {
	if r.err != nil {
		return r
	}
	r.err = err

	if r.sql != "" && sql != "" {
		r.sql += " "
	}
	r.sql += sql

	if r.bindings == nil {
		r.bindings = []any{}
	}
	if bindings != nil {
		r.bindings = append(r.bindings, bindings...)
	}
	return r
}

func (r *SQLResult) ToSQL(d dialects.Dialect) (string, []any, error) {
	return r.sql, r.bindings, r.err
}
