package pathutil

import (
	"os"
)

// Remove file
//	err := path.Remove()
// like os.Remove
func (path PathImpl) Remove() error {
	return os.Remove(path.path)
}

// Remove tree of directory(ies) include files
//  err := path.RemoveTree
// like os.RemoveAll
func (path PathImpl) RemoveTree() error {
	return os.RemoveAll(path.Canonpath())
}
