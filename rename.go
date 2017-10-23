package pathutil

import (
	"os"
)

// Rename path to new path
func (old *Path) Rename(new string) (*Path, error) {
	newPath, err := NewPath(new)
	if err != nil {
		return nil, err
	}

	err = os.Rename(old.Canonpath(), newPath.Canonpath())

	return newPath, err
}
