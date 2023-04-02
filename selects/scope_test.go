package selects_test

import (
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
)

func TestScope(t *testing.T) {
	scopeA := &selects.Scope{
		Name: "with-a",
		Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
			return b.Where("a", "=", "b")
		},
	}
	test.QueryTest(t, []test.Case{
		{
			Name:             "scope",
			Builder:          NewTestBuilder().WithScope(scopeA),
			ExpectedSQL:      "SELECT * FROM \"foos\" WHERE \"a\" = ?",
			ExpectedBindings: []any{"b"},
		},
		{
			Name:             "without scope",
			Builder:          NewTestBuilder().WithScope(scopeA).WithoutScope(scopeA),
			ExpectedSQL:      "SELECT * FROM \"foos\"",
			ExpectedBindings: []any{},
		},
	})
}
