package pathutil

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// NewPath construct *Path
//
// for example
//		path := NewPath("/home/test", ".vimrc")
//
func NewPath(path ...string) (*Path, error) {
	newPath := new(Path)

	joinPath := filepath.Join(path...)
	if len(joinPath) == 0 {
		return nil, errors.New("Paths requires defined, positive-lengths parts")
	}

	newPath.Path = strings.Replace(filepath.Clean(joinPath), "\\", "/", -1)

	return newPath, nil
}

// NewTempFile create temp file
//
// for cleanup use defer
//		temp, err := NewTempFile(TempOpt{})
//		defer temp.Remove()
//
// if you need only temp file name, you must delete file
//		temp, err := NewTempFile(TempFileOpt{})
//		temp.Remove()
//

func NewTempFile(options TempOpt) (*Path, error) {
	file, err := ioutil.TempFile(options.Dir, options.Prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempFile(%+v) fail", options)
	}

	defer file.Close()

	return NewPath(file.Name())
}

// NewTempDir create temp directory
//
// for cleanup use `defer`
// 	tempdir, err := pathutil.NewTempDir(TempOpt{})
//  defer tempdir.RemoveTree()
func NewTempDir(options TempOpt) (*Path, error) {
	dir, err := ioutil.TempDir(options.Dir, options.Prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempDir(%+v) fail", options)
	}

	return NewPath(dir)
}
