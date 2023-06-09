package selects_test

import (
	"context"
	"testing"

	"github.com/abibby/bob/selects"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	q := NewTestBuilder()

	assert.Equal(t, context.Background(), q.Context())

	q = q.WithContext(context.WithValue(context.Background(), "foo", "bar")).Where("1", "=", 2)

	assert.NotEqual(t, context.Background(), q.Context())
	assert.NotNil(t, q.Context())
}

func TestContext_sub_builder(t *testing.T) {
	q := NewTestBuilder()

	assert.Equal(t, context.Background(), q.Context())

	q = q.WithContext(context.WithValue(context.Background(), "foo", "bar")).Where("1", "=", 2)

	q.WhereHas("Bar", func(q *selects.SubBuilder) *selects.SubBuilder {
		assert.NotEqual(t, context.Background(), q.Context())
		assert.NotNil(t, q.Context())
		return q
	})
}
