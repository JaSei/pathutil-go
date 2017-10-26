package pathutil

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCopyFileDstFileExists(t *testing.T) {
	src, _ := NewPath("copy_test.go")
	dst, err := NewTempFile(TempOpt{})
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
