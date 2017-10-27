package pathutil

import (
	"os"
)

// Make path create directory(ies) in path if not exists (like `mkdir -p`) with default 0777 mode
// if you need set mode, use `MakePathMode`
func (path pathImpl) MakePath() error {
	return path.MakePathMode(0777)
}

// Make path create directory(ies) in path if not exists (like `mkdir -p`) with default given mode
func (path pathImpl) MakePathMode(mode os.FileMode) error {
	return os.MkdirAll(path.Canonpath(), mode)
}
