package pathutil

import (
	"fmt"
	"os"
)

// Stat return os.FileInfo
func (path pathImpl) Stat() (os.FileInfo, error) {
	file, err := os.Open(path.Path)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return file.Stat()
}

// File or dir exists

func (path pathImpl) Exists() bool {
	if _, err := path.Stat(); os.IsNotExist(err) {
		return false
	}

	return true
}

// IsDir return true if path is dir
func (path pathImpl) IsDir() bool {
	stat, err := path.Stat()
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// IsFile return true is path exists and not dir
// (symlinks, devs, regular files)
func (path pathImpl) IsFile() bool {
	return path.Exists() && !path.IsDir()
}

// IsRegularFile return true if path is regular file
// (wihtout devs, symlinks, ...)
func (path pathImpl) IsRegularFile() bool {
	stat, err := path.Stat()
	if err != nil {
		return false
	}

	return stat.Mode().IsRegular()
}
