package fields

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContextKey(t *testing.T) {
	expected_value := "fields context"

	ctx := context.WithValue(context.Background(),
		ContextKey, "fields context")

	assert.Equal(t, ctx.Value(ContextKey), expected_value)
}
