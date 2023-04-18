package selects

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
)

type limit struct {
	limit  int
	offset int
}

func (l *limit) Clone() *limit {
	return &limit{
		limit:  l.limit,
		offset: l.offset,
	}
}
func (l *limit) ToSQL(d dialects.Dialect) (string, []any, error) {
	if l.limit == 0 && l.offset == 0 {
		return "", nil, nil
	}
	r := builder.Result()
	r.AddString("LIMIT")
	r.Add(builder.Literal(l.limit))
	if l.offset != 0 {
		r.AddString("OFFSET")
		r.Add(builder.Literal(l.offset))
	}
	return r.ToSQL(d)
}

func (l *limit) Limit(limit int) *limit {
	l.limit = limit
	return l
}
func (l *limit) Offset(offset int) *limit {
	l.offset = offset
	return l
}
