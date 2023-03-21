package selects

import (
	"fmt"

	"github.com/abibby/bob/dialects"
)

type Limit struct {
	limit  int
	offset int
}

func (l *Limit) ToSQL(d dialects.Dialect) (string, []any, error) {
	if l.limit == 0 && l.offset == 0 {
		return "", nil, nil
	}
	r := &sqlResult{}
	r.addString(fmt.Sprintf("LIMIT %d", l.limit))
	if l.offset != 0 {
		r.addString(fmt.Sprintf("OFFSET %d", l.offset))
	}
	return r.ToSQL(d)
}

func (b *Builder) Limit(limit int) *Builder {
	b.limit.limit = limit
	return b
}
func (b *Builder) Offset(offset int) *Builder {
	b.limit.offset = offset
	return b
}
