package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent(t *testing.T) {
	temp, err := NewTempDir(TempOpt{})
	assert.NoError(t, err)

	assert.Equal(t, "/tmp", temp.Parent().String())

	root, err := New("/")
	assert.NoError(t, err)
	assert.Equal(t, "/", root.Parent().String())
}
