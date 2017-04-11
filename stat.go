package path

import (
	"fmt"
	"os"
)

func (path *Path) Stat() (os.FileInfo, error) {
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

func (path *Path) Exists() bool {
	if _, err := path.Stat(); os.IsNotExist(err) {
		return false
	}

	return true
}

// IsDir return true if path is dir
func (path *Path) IsDir() (bool, error) {
	stat, err := path.Stat()
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}

func (path *Path) IsFile() (bool, error) {
	stat, err := path.Stat()
	if err != nil {
		return false, err
	}

	return stat.Mode().IsRegular(), nil
}
