package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChdir(t *testing.T) {
	tempdir, err := NewTempDir(TempOpt{})
	assert.NoError(t, err)

	cwd, err := Cwd()
	assert.NoError(t, err)

	assert.NotEqual(t, tempdir, cwd, "Current working directory isn't same as tempdir")

	oldCwd, err := tempdir.Chdir()
	assert.NoError(t, err)

	// return cwd back
	defer func() {
		_, err = oldCwd.Chdir()
		assert.NoError(t, err)
	}()

	assert.NotEqual(t, tempdir, cwd, "Old current working directory isn't same as tempdir")
	assert.Equal(t, oldCwd, cwd, "Old current working directory is same as returned oldCwd after Chdir")

	actualCwd, err := Cwd()
	assert.NoError(t, err)

	assert.Equal(t, tempdir, actualCwd, "Actual current working directory is tempdir")
}
