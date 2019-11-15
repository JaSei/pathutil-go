package pathutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent(t *testing.T) {
	temp, err := NewTempDir()
	assert.NoError(t, err)

	tempDir, err := New(os.TempDir())
	assert.NoError(t, err)

	assert.Equal(t, tempDir.String(), temp.Parent().String())

	root, err := New("/")
	assert.NoError(t, err)
	assert.Equal(t, "/", root.Parent().String())
}
