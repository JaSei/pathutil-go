package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePath(t *testing.T) {
	tempdir, err := NewTempDir(TempOpt{})
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, tempdir.RemoveTree())
	}()

	newPath, err := NewPath(tempdir.String(), "a/b/c")
	assert.NoError(t, err)

	assert.False(t, newPath.Exists())

	assert.NoError(t, newPath.MakePath())

	assert.True(t, newPath.IsDir())
}
