package path

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// NewPath construct *Path
//
// for example
//		path := NewPath("/home/test", ".vimrc")
//
func NewPath(path ...string) (*Path, error) {
	newPath := new(Path)
	newPath.Path = strings.Replace(filepath.Join(path...), "\\", "/", -1)

	if len(newPath.Path) == 0 {
		return nil, errors.New("Paths requires defined, positive-lengths parts")
	}

	return newPath, nil
}

// NewTempFile create temp file
//
// for cleanup use defer
//		temp, err := NewTempFile(TempFileOpt{})
//		defer temp.Remove()

func NewTempFile(options TempFileOpt) (*Path, error) {
	dir := options.Dir

	if dir == "" {
		dir = os.TempDir()
	}

	file, err := ioutil.TempFile(dir, options.Prefix)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return &Path{
		file: file,
		Path: file.Name(),
	}, nil
}
