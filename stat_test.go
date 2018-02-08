package pathutil

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	path, _ := New("stat_test.go")
	assert.Exactly(t, true, path.Exists(), "file exists")
	path, _ = New("/sdlkfjsflsjfsl")
	assert.Exactly(t, false, path.Exists(), "wired root dir don't exists")
	path, _ = New(os.TempDir())
	assert.Exactly(t, true, path.Exists(), "home dir exists")
}

func TestIsDir(t *testing.T) {
	path, _ := New(os.TempDir())
	assert.Exactly(t, true, path.IsDir(), "temp dir is dir")

	path, _ = New("stat_test.go")
	assert.Exactly(t, false, path.IsDir(), "this test file isn't dir")

	path, _ = New("/safjasfjalfja")
	assert.Exactly(t, false, path.IsDir(), "unexists somethings isn't dir")
}

func TestIsFile(t *testing.T) {
	path, _ := New(os.TempDir())
	assert.Exactly(t, false, path.IsFile(), "temp dir is dir - no file")

	path, _ = New("stat_test.go")
	assert.Exactly(t, true, path.IsFile(), "this test file is file")

	path, _ = New("/safjasfjalfja")
	assert.Exactly(t, false, path.IsFile(), "unexists something isn't file")

	if runtime.GOOS != "windows" {
		path, _ = New("/dev/zero")
		assert.Exactly(t, true, path.IsFile(), "/dev/zero is file")
	}

	//symlink test
}

func TestIsRegularFile(t *testing.T) {
	path, _ := New(os.TempDir())
	assert.Exactly(t, false, path.IsRegularFile(), "temp dir is dir - no file")

	path, _ = New("stat_test.go")
	assert.Exactly(t, true, path.IsRegularFile(), "this test file is file")

	path, _ = New("/safjasfjalfja")
	assert.Exactly(t, false, path.IsRegularFile(), "unexists something isn't file")

	path, _ = New("/dev/zero")
	assert.Exactly(t, false, path.IsRegularFile(), "/dev/zero isn't regular file file")

	//symlink test
}
