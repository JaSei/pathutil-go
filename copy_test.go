package path

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestString(t *testing.T) {
	path, _ := NewPath("/tmp/test")
	assert.Equal(t, "/tmp/test", path.String(), "string")

	path, _ = NewPath("/tmp/test/")
	assert.Equal(t, "/tmp/test", path.String(), "string")

	path, _ = NewPath("/tmp", "test")
	assert.Equal(t, "/tmp/test", path.String(), "string")
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

func TestTempFile(t *testing.T) {
	temp, err := NewTempFile(TempFileOpt{Prefix: "bla"})
	defer temp.Remove()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}

func TestCopyFileDstFileNotExists(t *testing.T) {
	src, err := NewPath("copy_test.go")
	assert.Nil(t, err)

	dst, err := NewTempFile(TempFileOpt{})
	assert.Nil(t, err)
	defer dst.Remove()
	dst.Remove()

	assert.Exactly(t, false, dst.Exists(), "dst file isn't exists")
	newDst, err := src.CopyFile(dst.String())
	assert.Nil(t, err)
	if err != nil {
		t.Log(err)
	} else {
		assert.Equal(t, dst.String(), newDst.String(), "dst file is same as what return CopyFile")
		assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
	}
}

func TestCopyFileDstFileExists(t *testing.T) {
	src, _ := NewPath("copy_test.go")
	dst, err := NewTempFile(TempFileOpt{})
	assert.Nil(t, err)
	defer dst.Remove()

	assert.Exactly(t, true, dst.Exists(), "dst file exists before copy")
	newDst, err := src.CopyFile(dst.String())
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, dst.String(), newDst.String(), "dst file is same as what return CopyFile")
	assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
}

func TestCopyFileDstIsDir(t *testing.T) {
	src, _ := NewPath("copy_test.go")
	dst, _ := NewPath(os.TempDir())

	assert.Exactly(t, true, dst.Exists(), "dst dir exists")
	newDst, err := src.CopyFile(dst.String())
	if !assert.Nil(t, err) {
		return
	}

	defer newDst.Remove()

	assert.Exactly(t, true, newDst.Exists(), "dst after copy exists")
	assert.Exactly(t, true, newDst.IsFile(), "dst is file")
	assert.Equal(t, src.Basename(), newDst.Basename(), "if is dst directory, then is used basename of source")
}
