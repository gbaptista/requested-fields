package fields

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortbyLength(t *testing.T) {
	initial_slice := []string{"id", "product", "name"}

	expected_slice := []string{"product", "name", "id"}

	assert.NotEqual(t, expected_slice, initial_slice)

	sort.Sort(byLength(initial_slice))

	assert.Equal(t, expected_slice, initial_slice)
}
