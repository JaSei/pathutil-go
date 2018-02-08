package pathutil

import (
	"path/filepath"
)

// String return stable string representation of path
// this representation is linux like (slash as separator)
// for os specific string use Canonpath method
func (path PathImpl) String() string {
	return path.path
}

// Canonpath retrun path with right os separator
func (path PathImpl) Canonpath() string {
	return filepath.FromSlash(filepath.Clean(path.path))
}

// Basename
// like path/filepath.Base
func (path PathImpl) Basename() string {
	if path.path == "/" {
		return ""
	}

	return filepath.Base(path.path)
}
