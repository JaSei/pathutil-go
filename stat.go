package pathutil

import (
	"fmt"
	"os"
)

// Stat return os.FileInfo
func (path PathImpl) Stat() (os.FileInfo, error) {
	file, err := os.Open(path.path)

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

func (path PathImpl) Exists() bool {
	if _, err := path.Stat(); os.IsNotExist(err) {
		return false
	}

	return true
}

// IsDir return true if path is dir
func (path PathImpl) IsDir() bool {
	stat, err := path.Stat()
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// IsFile return true is path exists and not dir
// (symlinks, devs, regular files)
func (path PathImpl) IsFile() bool {
	return path.Exists() && !path.IsDir()
}

// IsRegularFile return true if path is regular file
// (wihtout devs, symlinks, ...)
func (path PathImpl) IsRegularFile() bool {
	stat, err := path.Stat()
	if err != nil {
		return false
	}

	return stat.Mode().IsRegular()
}
