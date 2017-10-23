package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {
	a, err := NewTempFile(TempFileOpt{})
	assert.NoError(t, err)

	b, err := NewTempFile(TempFileOpt{})
	assert.NoError(t, err)
	assert.NoError(t, b.Remove())

	assert.True(t, a.Exists(), "path a exists")
	assert.False(t, b.Exists(), "path b don't")

	c, err := a.Rename(b.Canonpath())
	assert.NoError(t, err)

	assert.False(t, a.Exists(), "After raname path a not exists")
	assert.True(t, b.Exists(), "After raname path b exists")

	assert.Equal(t, b, c, "Path b and c is same")
}
