package path

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestString(t *testing.T) {
	assert.Equal(t, "/tmp/test", NewPath("/tmp/test").String(), "string")
	assert.Equal(t, "/tmp/test", NewPath("/tmp/test/").String(), "string")

	assert.Equal(t, "/tmp/test", NewPath("/tmp", "test").String(), "string")
	assert.Equal(t, "/tmp/test", NewPath("/tmp", NewPath("test").String()).String(), "string")
}

func TestBasename(t *testing.T) {
	assert.Equal(t, "test", NewPath("/tmp/test").Basename(), "basename of /tmp/test")
	assert.Equal(t, "", NewPath("/").Basename(), "basename of root")
}

func TestExists(t *testing.T) {
	assert.Exactly(t, true, NewPath("copy_test.go").Exists(), "this test exists")
	assert.Exactly(t, false, NewPath("/sdlkfjsflsjfsl").Exists(), "wired root dir don't exists")
	assert.Exactly(t, true, NewPath(os.TempDir()).Exists(), "home dir exists")
}

func TestTempFile(t *testing.T) {
	temp, err := NewTempFile(TempFileOpt{Prefix: "bla"})
	defer temp.Remove()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}

func TestIsDir(t *testing.T) {
	isDir, err := NewPath(os.TempDir()).IsDir()
	assert.Nil(t, err)
	assert.Exactly(t, true, isDir, "temp dir is dir")

	isDir, err = NewPath("copy_test.go").IsDir()
	assert.Nil(t, err)
	assert.Exactly(t, false, isDir, "this test file isn't dir")

	isDir, err = NewPath("/safjasfjalfja").IsDir()
	assert.NotNil(t, err)
	assert.Exactly(t, false, isDir, "unexists somethings isn't dir")
}

func TestIsFile(t *testing.T) {
	isFile, err := NewPath(os.TempDir()).IsFile()
	assert.Nil(t, err)
	assert.Exactly(t, false, isFile, "temp dir is dir - no file")

	isFile, err = NewPath("copy_test.go").IsFile()
	assert.Nil(t, err)
	assert.Exactly(t, true, isFile, "this test file is file")

	isFile, err = NewPath("/safjasfjalfja").IsFile()
	assert.NotNil(t, err)
	assert.Exactly(t, false, isFile, "unexists somethings isn't file")
}

func TestCopyFile(t *testing.T) {
	t.Run("dst file not exists",
		func(t *testing.T) {
			src := NewPath("copy_test.go")
			dst, err := NewTempFile(TempFileOpt{})
			assert.Nil(t, err)
			defer dst.Remove()
			dst.Remove()

			assert.Exactly(t, false, dst.Exists(), "dst file isn't exists")
			newDst, err := src.CopyFile(dst.String())
			assert.Nil(t, err)
			assert.Equal(t, dst.String(), newDst.String(), "dst file is same as what return CopyFile")
			assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
		})

	t.Run("dst file exists",
		func(t *testing.T) {
			src := NewPath("copy_test.go")
			dst, err := NewTempFile(TempFileOpt{})
			assert.Nil(t, err)
			defer dst.Remove()

			assert.Exactly(t, true, dst.Exists(), "dst file exists before copy")
			newDst, err := src.CopyFile(dst.String())
			assert.Nil(t, err)
			assert.Equal(t, dst.String(), newDst.String(), "dst file is same as what return CopyFile")
			assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
		})

	t.Run("dst is dir",
		func(t *testing.T) {
			src := NewPath("copy_test.go")
			dst := NewPath(os.TempDir())

			assert.Exactly(t, true, dst.Exists(), "dst dir exists")
			newDst, err := src.CopyFile(dst.String())
			defer newDst.Remove()

			assert.Nil(t, err)
			assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
			assert.Equal(t, src.Basename(), newDst.Basename(), "if is dst directory, then is used basename of source")
		})

}
