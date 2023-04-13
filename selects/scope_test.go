package selects_test

import (
	"context"
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

type ScopeBar struct {
	test.Bar
	ScopeFoo *selects.BelongsTo[*ScopeFoo] `db:"-" json:"foo"`
}

func TestScope(t *testing.T) {
	scopeA := &selects.Scope{
		Name: "with-a",
		Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
			return b.Where("a", "=", "b")
		},
	}
	scopeCtx := &selects.Scope{
		Name: "ctx",
		Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
			foo := b.Context().Value("foo")
			return b.Where("a", "=", foo)
		},
	}
	test.QueryTest(t, []test.Case{
		{
			Name:             "scope",
			Builder:          NewTestBuilder().WithScope(scopeA),
			ExpectedSQL:      "SELECT \"foos\".* FROM \"foos\" WHERE \"a\" = ?",
			ExpectedBindings: []any{"b"},
		},
		{
			Name:             "without scope",
			Builder:          NewTestBuilder().WithScope(scopeA).WithoutScope(scopeA),
			ExpectedSQL:      "SELECT \"foos\".* FROM \"foos\"",
			ExpectedBindings: []any{},
		},
		{
			Name:             "global scope",
			Builder:          selects.From[*ScopeFoo](),
			ExpectedSQL:      "SELECT \"scope_foos\".* FROM \"scope_foos\" WHERE \"deleted_at\" IS NULL",
			ExpectedBindings: []any{},
		},
		{
			Name:             "without global scope",
			Builder:          selects.From[*ScopeFoo]().WithoutGlobalScope(bob.SoftDeletes),
			ExpectedSQL:      "SELECT \"scope_foos\".* FROM \"scope_foos\"",
			ExpectedBindings: []any{},
		},
		{
			Name: "global scope whereHas",
			Builder: selects.From[*ScopeBar]().WhereHas("ScopeFoo", func(q *selects.SubBuilder) *selects.SubBuilder {
				return q
			}),
			ExpectedSQL:      `SELECT "scope_bars".* FROM "scope_bars" WHERE EXISTS (SELECT "scope_foos".* FROM "scope_foos" WHERE "id" = "scope_bars"."scope_foo_id" AND "deleted_at" IS NULL)`,
			ExpectedBindings: []any{},
		},
		{
			Name:             "access-context",
			Builder:          NewTestBuilder().WithScope(scopeCtx).WithContext(context.WithValue(context.Background(), "foo", "bar")),
			ExpectedSQL:      "SELECT \"foos\".* FROM \"foos\" WHERE \"a\" = ?",
			ExpectedBindings: []any{"bar"},
		},
	})
}
