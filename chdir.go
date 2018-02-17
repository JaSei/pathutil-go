package pathutil

import (
	"os"

	"github.com/pkg/errors"
)

// Chdir change current working directory do the path and return old current working directory
func (path PathImpl) Chdir() (Path, error) {
	oldPath, err := Cwd()
	if err != nil {
		return nil, errors.Wrap(err, "Cwd fail")
	}

	if err := os.Chdir(path.Canonpath()); err != nil {
		return nil, errors.Wrapf(err, "Chdir(%s) fail", path)
	}

	return oldPath, nil
}
