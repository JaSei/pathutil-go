package pathutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestString(t *testing.T) {
	path, _ := NewPath("tmp/test")
	assert.Equal(t, "tmp/test", path.String(), "string")

	path, _ = NewPath("tmp/test/")
	assert.Equal(t, "tmp/test", path.String(), "string")

	path, _ = NewPath("tmp", "test")
	assert.Equal(t, "tmp/test", path.String(), "string")
}

func TestCanonpath(t *testing.T) {
	path, _ := NewPath("tmp/test")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "linux rel path")

	path, _ = NewPath("tmp/test/")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "linux rel path with end slash")

	path, _ = NewPath("tmp", "test")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "more paths elemets")

	path, _ = NewPath("tmp\\test")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "windows rel path")
}

func TestBasename(t *testing.T) {
	path, _ := NewPath("/tmp/test")
	assert.Equal(t, "test", path.Basename(), "basename of /tmp/test")

	path, _ = NewPath("/")
	assert.Equal(t, "", path.Basename(), "basename of root")

	path, _ = NewPath("relative")
	assert.Equal(t, "relative", path.Basename(), "relative basename")

	path, _ = NewPath("relative/a")
	assert.Equal(t, "a", path.Basename(), "relative subpath basename")
}
