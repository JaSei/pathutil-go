package path

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	path, _ := NewPath("copy_test.go")
	assert.Exactly(t, true, path.Exists(), "this test exists")
	path, _ = NewPath("/sdlkfjsflsjfsl")
	assert.Exactly(t, false, path.Exists(), "wired root dir don't exists")
	path, _ = NewPath(os.TempDir())
	assert.Exactly(t, true, path.Exists(), "home dir exists")
}

func TestIsDir(t *testing.T) {
	path, _ := NewPath(os.TempDir())
	assert.Exactly(t, true, path.IsDir(), "temp dir is dir")

	path, _ = NewPath("copy_test.go")
	assert.Exactly(t, false, path.IsDir(), "this test file isn't dir")

	path, _ = NewPath("/safjasfjalfja")
	assert.Exactly(t, false, path.IsDir(), "unexists somethings isn't dir")
}

func TestIsFile(t *testing.T) {
	path, _ := NewPath(os.TempDir())
	assert.Exactly(t, false, path.IsFile(), "temp dir is dir - no file")

	path, _ = NewPath("copy_test.go")
	assert.Exactly(t, true, path.IsFile(), "this test file is file")

	path, _ = NewPath("/safjasfjalfja")
	assert.Exactly(t, false, path.IsFile(), "unexists somethings isn't file")

	path, _ = NewPath("/dev/zero")
	assert.Exactly(t, true, path.IsFile(), "/dev/zero is file")

	//symlink test
}

func TestIsRegularFile(t *testing.T) {
	path, _ := NewPath(os.TempDir())
	assert.Exactly(t, false, path.IsRegularFile(), "temp dir is dir - no file")

	path, _ = NewPath("copy_test.go")
	assert.Exactly(t, true, path.IsRegularFile(), "this test file is file")

	path, _ = NewPath("/safjasfjalfja")
	assert.Exactly(t, false, path.IsRegularFile(), "unexists somethings isn't file")

	path, _ = NewPath("/dev/zero")
	assert.Exactly(t, false, path.IsRegularFile(), "/dev/zero is file")

	//symlink test
}
