package pathutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
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

// TempDir set directory where is temp file/dir create, empty string `""` (default) means TEMPDIR (`os.TempDir`)
func TempDir(dir string) TempOpt {
	return func(o *tempOpt) {
		o.dir = dir
	}
}

// TempPrefix set name beginning with prefix
func TempPrefix(prefix string) TempOpt {
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
//		temp, err := NewTempFile()
//		temp.Remove()
//
// if you need set directory or prefix, then use `TempDir` and/or `TempPrefix`
//		temp, err := NewTempFile(TempDir("/home/my/test"), TempPrefix("myfile"))
//		...
//
func NewTempFile(options ...TempOpt) (p Path, err error) {
	opt := tempOpt{}
	for _, o := range options {
		o(&opt)
	}

	file, err := tempFile(opt.dir, opt.prefix)
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
// 	tempdir, err := pathutil.NewTempDir()
//  defer tempdir.RemoveTree()
//
// if you need set directory or prefix, then use `TempDir` and/or `TempPrefix`
//		temp, err := NewTempFile(TempDir("/home/my/test"), TempPrefix("myfile"))
//		...
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

// Cwd create new Path from current working directory
// optional is possible to set subpath
//
// for example
//		gitConfigPath, err := pathutil.Cwd('.git/config')
//
func Cwd(subpath ...string) (Path, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "Getwd fail")
	}

	return New(join(cwd, subpath)...)
}

// Home create new Path from home directory
// optional is possible to set subpath
//
// for example
//		initPath, err := pathutil.Home('.config/nvim/init.vim')
//
// (internally use https://github.com/mitchellh/go-homedir library)
func Home(subpath ...string) (Path, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, errors.Wrap(err, "homedir.Dir fail")
	}

	return New(join(home, subpath)...)
}

func join(a string, b []string) []interface{} {
	p := make([]interface{}, len(b)+1)
	p[0] = a
	for i, c := range b {
		p[i+1] = c
	}

	return p
}
