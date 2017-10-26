package pathutil

import (
	"os"
)

// Remove file
//	err := path.Remove()
// like os.Remove
func (path *Path) Remove() error {
	return os.Remove(path.Path)
}

// Remove tree of directory(ies) include files
//  err := path.RemoveTree
// like os.RemoveAll
func (path *Path) RemoveTree() error {
	return os.RemoveAll(path.Canonpath())
}
