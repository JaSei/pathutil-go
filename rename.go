package pathutil

import (
	"os"
)

// Rename path to new path
func (old PathImpl) Rename(new string) (Path, error) {
	newPath, err := New(new)
	if err != nil {
		return nil, err
	}

	err = os.Rename(old.Canonpath(), newPath.Canonpath())

	return newPath, err
}
