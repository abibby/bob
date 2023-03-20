package bob

import (
	"github.com/abibby/bob/dialects"
)

type ToSQLer interface {
	ToSQL(d dialects.Dialect) (string, []any, error)
}

type sqlResult struct {
	sql  string
	args []any
	err  error
}

func (r *sqlResult) addString(sql string) {
	r.add(sql, nil, nil)
}
func (r *sqlResult) add(sql string, args []any, err error) {
	if r.err != nil {
		return
	}
	r.err = err

	if r.sql != "" && sql != "" {
		r.sql += " "
	}
	r.sql += sql

	if r.args == nil {
		r.args = []any{}
	}
	if args != nil {
		r.args = append(r.args, args...)
	}
}

func (r *sqlResult) ToSQL(d dialects.Dialect) (string, []any, error) {
	return r.sql, r.args, r.err
}

func (b *Builder) ToSQL(d dialects.Dialect) (string, []any, error) {
	r := &sqlResult{}
	r.add(b.selects.ToSQL(d))
	r.add(b.from.ToSQL(d))
	r.add(b.wheres.ToSQL(d))
	r.add(b.groupBys.ToSQL(d))
	r.add(b.havings.ToSQL(d))

	return r.ToSQL(d)
}
