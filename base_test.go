package pathutil

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	path, _ := New("tmp/test")
	assert.Equal(t, "tmp/test", path.String(), "string")

	path, _ = New("tmp/test/")
	assert.Equal(t, "tmp/test", path.String(), "string")

	path, _ = New("tmp", "test")
	assert.Equal(t, "tmp/test", path.String(), "string")
}

func TestCanonpath(t *testing.T) {
	path, _ := New("tmp/test")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "linux rel path")

	path, _ = New("tmp/test/")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "linux rel path with end slash")

	path, _ = New("tmp", "test")
	assert.Equal(t, fmt.Sprintf("tmp%stest", string(filepath.Separator)), path.Canonpath(), "more paths elemets")

	path, _ = New("tmp\\test")
	assert.Equal(t, "tmp\\test", path.Canonpath(), "windows rel path")
	if runtime.GOOS == "windows" {
		assert.Equal(t, "tmp/test", path.String(), "windows backshlash are internaly represents as slash")
	} else {
		assert.Equal(t, "tmp\\test", path.String(), "backshlash in linux path is allowed")
	}
}

func TestBasename(t *testing.T) {
	path, _ := New("/tmp/test")
	assert.Equal(t, "test", path.Basename(), "basename of /tmp/test")

	path, _ = New("/")
	assert.Equal(t, "", path.Basename(), "basename of root")

	path, _ = New("relative")
	assert.Equal(t, "relative", path.Basename(), "relative basename")

	path, _ = New("relative/a")
	assert.Equal(t, "a", path.Basename(), "relative subpath basename")
}
