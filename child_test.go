package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChild(t *testing.T) {
	tempdir, err := NewTempDir()
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, tempdir.RemoveTree())
	}()

	kukPath, err := tempdir.Child("KUK", "PUK")
	assert.NoError(t, err)
	assert.NoError(t, kukPath.MakePath())

	assert.True(t, kukPath.Exists())
}

func TestChildren(t *testing.T) {
	tempdir, err := NewTempDir()
	assert.NoError(t, err)

	defer func() {
		assert.NoError(t, tempdir.RemoveTree())
	}()

	a, _ := tempdir.Child("a", "c")
	assert.NoError(t, a.MakePath())
	b, _ := tempdir.Child("b")
	assert.NoError(t, b.MakePath())

	children, err := tempdir.Children()
	assert.NoError(t, err)

	assert.Len(t, children, 2)

	exA, _ := New(tempdir.String(), "a")
	exB, _ := New(tempdir.String(), "b")
	assert.Equal(t, []Path{exA, exB}, children)
}
