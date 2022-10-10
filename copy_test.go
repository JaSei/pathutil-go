package pathutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFileDstFileExists(t *testing.T) {
	src, _ := New("copy_test.go")
	dst, err := NewTempFile()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, dst.Remove())
	}()

	assert.Exactly(t, true, dst.Exists(), "dst file exists before copy")
	newDst, err := src.CopyFile(dst.String())
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, dst.String(), newDst.String(), "dst file is same as what return CopyFile")
	assert.Exactly(t, true, newDst.Exists(), "dst file after copy exists")
}

func TestCopyFileDstIsDir(t *testing.T) {
	src, _ := New("copy_test.go")
	dst, _ := New(os.TempDir())

	assert.Exactly(t, true, dst.Exists(), "dst dir exists")
	newDst, err := src.CopyFile(dst.String())
	if !assert.NoError(t, err) {
		return
	}

	defer func() {
		assert.NoError(t, newDst.Remove())
	}()

	assert.Exactly(t, true, newDst.Exists(), "dst after copy exists")
	assert.Exactly(t, true, newDst.IsFile(), "dst is file")
	assert.Equal(t, src.Basename(), newDst.Basename(), "if is dst directory, then is used basename of source")
}
