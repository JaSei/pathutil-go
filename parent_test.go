package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParent(t *testing.T) {
	temp, err := NewTempDir(TempOpt{})
	assert.NoError(t, err)

	parent, err := temp.Parent()
	assert.NoError(t, err)
	assert.Equal(t, "/tmp", parent.String())

	root, err := NewPath("/")
	assert.NoError(t, err)
	parentOfRoot, err := root.Parent()
	assert.NoError(t, err)
	assert.Equal(t, "/", parentOfRoot.String())
}
