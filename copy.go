package path

import (
	"io"
	"os"
	"path/filepath"
)

// Remove file
//	err := path.Remove()
// like os.Remove

func (path *Path) Remove() error {
	path.file = nil
	return os.Remove(path.Path)
}

// Remove tree of files
//  err := path.RemoveTree
// like os.RemoveAll
func (path *Path) RemoveTree() error {
	path.file = nil
	return os.RemoveAll(path.Path)
}

func (path *Path) String() string {
	return filepath.FromSlash(filepath.Clean(path.Path))
}

func (path *Path) Basename() string {
	if path.Path == "/" {
		return ""
	}

	return filepath.Base(path.Path)
}

func (srcPath *Path) CopyFile(dst string) (*Path, error) {
	dstPath, err := NewPath(dst)
	if err != nil {
		return nil, err
	}

	if dstPath.IsDir() {
		dstPath, err := NewPath(dst, srcPath.Basename())
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

	return NewPath(dst)
}
