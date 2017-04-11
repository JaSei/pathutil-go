package path

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewPath construct *Path
//
// for example
//		path := NewPath("/home/test", ".vimrc")
//
func NewPath(path ...string) *Path {
	newPath := new(Path)
	newPath.Path = filepath.Join(path...)

	return newPath
}

// NewTempFile create temp file
//
// for delete after scope use defer
//		temp, err := NewTempFile(TempFileOpt{})
//		defer temp.Remove()

func NewTempFile(options TempFileOpt) (*Path, error) {
	dir := options.Dir

	if dir == "" {
		dir = os.TempDir()
	}

	file, err := ioutil.TempFile(dir, options.Prefix)

	if err != nil {
		return nil, err
	}

	return &Path{
		file: file,
		Path: file.Name(),
	}, nil
}
