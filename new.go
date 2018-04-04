package pathutil

import (
	"fmt"
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
// input can be `string` or type implements `fmt.Stringer` interface
func New(path ...interface{}) (Path, error) {
	newPath := PathImpl{}

	paths := make([]string, len(path))
	for index, chunk := range path {
		var pathChunk string
		switch t := chunk.(type) {
		case string:
			pathChunk = chunk.(string)
		case fmt.Stringer:
			pathChunk = chunk.(fmt.Stringer).String()
		default:
			return nil, errors.Errorf("Chunk %d is type %t (allowed are string or fmt.Stringer)", index, t)
		}

		if len(pathChunk) == 0 {
			return nil, errors.Errorf("Paths requires defined, positive-lengths parts (part %d is empty", index)
		}

		paths[index] = pathChunk
	}

	joinPath := filepath.Join(paths...)

	newPath.path = strings.Replace(filepath.Clean(joinPath), "\\", "/", -1)

	return newPath, nil
}

type tempOpt struct {
	dir    string
	prefix string
}

//TempOpt is func for configure tempdir or tempfile
type TempOpt func(*tempOpt)

// directory where is temp file/dir create, empty string `""` (default) means TEMPDIR (`os.TempDir`)
func Dir(dir string) TempOpt {
	return func(o *tempOpt) {
		o.dir = dir
	}
}

// name beginning with prefix
func Prefix(prefix string) TempOpt {
	return func(o *tempOpt) {
		o.prefix = prefix
	}
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

func NewTempFile(options ...TempOpt) (p Path, err error) {
	opt := tempOpt{}
	for _, o := range options {
		o(&opt)
	}

	file, err := ioutil.TempFile(opt.dir, opt.prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempFile(%+v) fail", opt)
	}

	defer func() {
		if errClose := file.Close(); errClose != nil {
			err = errClose
		}
	}()

	return New(file.Name())
}

// NewTempDir create temp directory
//
// for cleanup use `defer`
// 	tempdir, err := pathutil.NewTempDir(TempOpt{})
//  defer tempdir.RemoveTree()
func NewTempDir(options ...TempOpt) (Path, error) {
	opt := tempOpt{}
	for _, o := range options {
		o(&opt)
	}

	dir, err := ioutil.TempDir(opt.dir, opt.prefix)
	if err != nil {
		return nil, errors.Wrapf(err, "NewTempDir(%+v) fail", opt)
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
