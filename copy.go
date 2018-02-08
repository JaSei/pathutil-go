package pathutil

import (
	"io"
	"os"
)

func (srcPath PathImpl) CopyFile(dst string) (Path, error) {
	dstPath, err := New(dst)
	if err != nil {
		return nil, err
	}

	if dstPath.IsDir() {
		dstPath, err := New(dst, srcPath.Basename())
		if err != nil {
			return nil, err
		} else {
			dst = dstPath.String()
		}
	}

	originalFile, err := os.Open(srcPath.String())
	if err != nil {
		return nil, err
	}
	defer originalFile.Close()

	newFile, err := os.Create(dst)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return nil, err
	}

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		return nil, err
	}

	return New(dst)
}
