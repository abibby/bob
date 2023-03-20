package bob

import (
	"testing"

	"github.com/abibby/bob/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

type TestCase[T any] struct {
	name string
	q    *SelectBuilder
	sql  string
	args []any
}

func QueryTest[T any](t *testing.T, testCases []TestCase[T]) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q, args, err := tc.q.ToSQL(&mysql.MySQL{})
			assert.NoError(t, err)

			assert.Equal(t, tc.sql, q)
			assert.Equal(t, tc.args, args)
		})
	}
}
