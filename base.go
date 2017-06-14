package pathutil

import (
	"os"
	"path/filepath"
)

// Remove file
//	err := path.Remove()
// like os.Remove
func (path *Path) Remove() error {
	return os.Remove(path.Path)
}

// Remove tree of files
//  err := path.RemoveTree
// like os.RemoveAll
func (path *Path) RemoveTree() error {
	return os.RemoveAll(path.Path)
}

// String return stable string representation of path
// this representation is linux like (slash as separator)
// for os specific string use Canonpath method
func (path *Path) String() string {
	return filepath.Clean(path.Path)
}

// Canonpath retrun path with right os separator
func (path *Path) Canonpath() string {
	return filepath.FromSlash(filepath.Clean(path.Path))
}

// Basename
// like path/filepath.Base
func (path *Path) Basename() string {
	if path.Path == "/" {
		return ""
	}

	return filepath.Base(path.Path)
}