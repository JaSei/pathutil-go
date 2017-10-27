package pathutil

import (
	"path/filepath"
)

// String return stable string representation of path
// this representation is linux like (slash as separator)
// for os specific string use Canonpath method
func (path pathImpl) String() string {
	return path.Path
}

// Canonpath retrun path with right os separator
func (path pathImpl) Canonpath() string {
	return filepath.FromSlash(filepath.Clean(path.Path))
}

// Basename
// like path/filepath.Base
func (path pathImpl) Basename() string {
	if path.Path == "/" {
		return ""
	}

	return filepath.Base(path.Path)
}
