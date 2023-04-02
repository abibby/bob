package selects_test

import (
	"testing"

	"github.com/abibby/bob"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
)

type ScopeFoo test.Foo

func (f *ScopeFoo) Scopes() []*selects.Scope {
	return []*selects.Scope{
		bob.SoftDeletes,
	}
}

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
		{
			Name:             "global scope",
			Builder:          selects.From[*ScopeFoo](),
			ExpectedSQL:      "SELECT * FROM \"scope_foos\" WHERE \"deleted_at\" IS NOT NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "global scope",
			Builder:          selects.From[*ScopeFoo]().WithoutGlobalScope(bob.SoftDeletes),
			ExpectedSQL:      "SELECT * FROM \"scope_foos\"",
			ExpectedBindings: []any{},
		},
	})
}
