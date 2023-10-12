package ini

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInSlice(t *testing.T) {
	ss := []string{"a", "b", "c"}
	assert.True(t, inSlice("a", ss))
	assert.False(t, inSlice("d", ss))
}
