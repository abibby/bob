package test

import (
	"testing"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	Name             string
	Builder          builder.ToSQLer
	ExpectedSQL      string
	ExpectedBindings []any
}

func QueryTest(t *testing.T, testCases []Case) {

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			q, args, err := tc.Builder.ToSQL(&mysql.MySQL{})
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedSQL, q)
			assert.Equal(t, tc.ExpectedBindings, args)
		})
	}
}
