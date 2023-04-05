package selects_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	q := NewTestBuilder()

	assert.Equal(t, context.Background(), q.Context())

	q = q.WithContext(context.WithValue(context.Background(), "foo", "bar"))

	assert.NotEqual(t, context.Background(), q.Context())
}
