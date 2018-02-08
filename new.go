package pathutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// New construct Path
//
// for example
//		path := New("/home/test", ".vimrc")
//
//
// if you can use `Path` in `New`, you must use `.String()` method
func New(path ...string) (Path, error) {
	newPath := PathImpl{}

	for index, pathChunk := range path {
		if len(pathChunk) == 0 {
			return nil, errors.Errorf("Paths requires defined, positive-lengths parts (part %d is empty", index)
		}
	}

	joinPath := filepath.Join(path...)

	newPath.path = strings.Replace(filepath.Clean(joinPath), "\\", "/", -1)

	return newPath, nil
}

//TempOpt is struct for configure new tempdir or tempfile
type TempOpt struct {
	// directory where is temp file/dir create, empty string `""` (default) means TEMPDIR (`os.TempDir`)
	Dir string
	// name beginning with prefix
	Prefix string
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

func NewTempFile(options TempOpt) (Path, error) {
	file, err := ioutil.TempFile(options.Dir, options.Prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempFile(%+v) fail", options)
	}

	defer file.Close()

	return New(file.Name())
}

// NewTempDir create temp directory
//
// for cleanup use `defer`
// 	tempdir, err := pathutil.NewTempDir(TempOpt{})
//  defer tempdir.RemoveTree()
func NewTempDir(options TempOpt) (Path, error) {
	dir, err := ioutil.TempDir(options.Dir, options.Prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempDir(%+v) fail", options)
	}

	return New(dir)
}

// Cwd create new path from current working directory
func Cwd() (Path, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "Getwd fail")
	}

	return New(cwd)
}
