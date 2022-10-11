# pathutil

[![Release](https://img.shields.io/github/release/JaSei/pathutil-go.svg?style=flat-square)](https://github.com/JaSei/pathutil-go/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
![GitHub Actions](https://github.com/JaSei/pathutil-go/actions/workflows/workflow.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/JaSei/pathutil-go?style=flat-square)](https://goreportcard.com/report/github.com/JaSei/pathutil-go)
[![GoDoc](https://godoc.org/github.com/JaSei/pathutil-go?status.svg&style=flat-square)](http://godoc.org/github.com/JaSei/pathutil-go)
[![codecov.io](https://codecov.io/github/JaSei/pathutil-go/coverage.svg?branch=master)](https://codecov.io/github/JaSei/pathutil-go?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/JaSei/pathutil-go/-/badge.svg)](https://sourcegraph.com/github.com/JaSei/pathutil-go?badge)

Pathutil is I/O utility primary inspired by David Golden's
[Path::Tiny](https://metacpan.org/pod/Path::Tiny). It is friendlier to use than
[path](https://golang.org/pkg/path/),
[filepath](https://golang.org/pkg/path/filepath/) and provides many of other
functions which isn't in core libraries (like `Copy` for example)

### SYNOPSIS

    import (
    	"fmt"
    	"github.com/JaSei/pathutil-go"
    )

    // creating pathutil objects
    dir, _ := pathutil.New("/tmp")
    foo, _ := pathutil.New("foo.txt")

    subdir, _ := dir.Child("foo")
    bar, _ := subdir.Child("bar.txt")

    // stringifies as cleaned up path
    file, _ := pathutil.New("./foo.txt")
    fmt.Println(file) // "foo.txt"

    // reading files
    guts, _ := file.Slurp()
    lines, _ := file.Lines()

    // writing files
    bar.Spew(data)

    // reading directories
    children, _ := dir.Children()
    for _, child := range children {
    }


### SEE ALSO

* [Path::Tiny](https://metacpan.org/pod/Path::Tiny) for Perl

* [better files](https://github.com/pathikrit/better-files) for Scala

* [pathlib](https://docs.python.org/3/library/pathlib.html) for python

BREAKING CHANGE 0.3.1 -> 1.0.0

1. `NewTempFile` or `NewTempDir` don't need TempOpt struct

    //0.3.1 default
    pathutil.NewTempFile(pathutil.TempOpt{})
    //0.3.1 custom
    pathutil.NewTempFile(pathutil.TempOpt{Dir: "/test", Prefix: "pre"})

    //1.0.0 default
    pathutil.NewTempFile()
    //1.0.0 custom
    pathutil.NewTempFile(Dir("/test"), Prefix("pre"))

2. `New` method parameter allowed `string` type or type implements
`fmt.Stringer` interface

    //0.3.1
    pathutil.New(otherPath.String(), "test")

    //1.0.0
    pathutil.New(otherPath, "test")

This shouldn't be breaking change, but if you use in some code variadic
parameter as input of `pathutil.New`, then can be problem

    //0.3.1
    func(p ...string) {
    	pathutil.New(p...)
    }("a", "b")

    //1.0.0
    func(p ...string) {
    	n := make([]interface{}, len(p))
    	for i, v := range p {
    		n[i] = v
    	}
    	pathutil.New(n...)
    }("a", "b")

3. There is new (more handfull) crypto API

    //0.3.1
    import (
    	"crypto"
    	"github.com/JaSei/pathutil-go"
    )
    ...

    hash, err := path.Crypto(crypto.SHA256)
    if err == nil {
    	fmt.Printf("%s\t%s\n", hash.HexSum(), path.String())
    }

    //1.0.0
    import (
    	"github.com/JaSei/pathutil-go"
    )
    ...

    hash, err := path.CryptoSha256()
    if err == nil {
    	fmt.Printf("%s\t%s\n", hash, path.String())
    }

This new crypto API return [hashutil](github.com/JaSei/hashutil-go) struct which
is more handfull for compare, transformation and next hash manipulation.

## Usage

#### type LinesFunc

```go
type LinesFunc func(string)
```


#### type Path

```go
type Path interface {
	String() string
	Canonpath() string
	Basename() string

	Chdir() (Path, error)
	Rename(string) (Path, error)

	Stat() (os.FileInfo, error)

	IsDir() bool
	Exists() bool
	IsFile() bool
	IsRegularFile() bool

	Remove() error
	RemoveTree() error

	Visit(VisitFunc, VisitOpt)
	CopyFile(string) (Path, error)

	CopyFrom(io.Reader) (int64, error)

	CryptoMd5() (hashutil.Md5, error)
	CryptoSha1() (hashutil.Sha1, error)
	CryptoSha256() (hashutil.Sha256, error)
	CryptoSha384() (hashutil.Sha384, error)
	CryptoSha512() (hashutil.Sha512, error)

	MakePath() error
	MakePathMode(os.FileMode) error

	OpenReader() (ReadSeekCloser, error)
	OpenWriter() (*os.File, error)
	OpenWriterAppend() (*os.File, error)

	Slurp() (string, error)
	SlurpBytes() ([]byte, error)

	Spew(string) error
	SpewBytes([]byte) error

	Lines() ([]string, error)
	LinesWalker(LinesFunc) error

	Child(...string) (Path, error)
	Children() ([]Path, error)

	Parent() Path

	Append(string) error
	AppendBytes([]byte) error
}
```


#### func  Cwd

```go
func Cwd(subpath ...string) (Path, error)
```
Cwd create new Path from current working directory optional is possible to set
subpath

for example

    gitConfigPath, err := pathutil.Cwd('.git/config')

#### func  Home

```go
func Home(subpath ...string) (Path, error)
```
Home create new Path from home directory optional is possible to set subpath

for example

    initPath, err := pathutil.Home('.config/nvim/init.vim')

(internally use https://github.com/mitchellh/go-homedir library)

#### func  New

```go
func New(path ...interface{}) (Path, error)
```
New construct Path

for example

    path := New("/home/test", ".vimrc")

input can be `string` or type implements `fmt.Stringer` interface

#### func  NewTempDir

```go
func NewTempDir(options ...TempOpt) (Path, error)
```
NewTempDir create temp directory

for cleanup use `defer`

    	tempdir, err := pathutil.NewTempDir()
     defer tempdir.RemoveTree()

if you need set directory or prefix, then use `TempDir` and/or `TempPrefix`

    temp, err := NewTempFile(TempDir("/home/my/test"), TempPrefix("myfile"))
    ...

#### func  NewTempFile

```go
func NewTempFile(options ...TempOpt) (p Path, err error)
```
NewTempFile create temp file

for cleanup use defer

    temp, err := NewTempFile(TempOpt{})
    defer temp.Remove()

if you need only temp file name, you must delete file

    temp, err := NewTempFile()
    temp.Remove()

if you need set directory or prefix, then use `TempDir` and/or `TempPrefix`

    temp, err := NewTempFile(TempDir("/home/my/test"), TempPrefix("myfile"))
    ...

#### type PathImpl

```go
type PathImpl struct {
}
```

type PathImpl implements Path interface

#### func (PathImpl) Append

```go
func (path PathImpl) Append(data string) error
```

#### func (PathImpl) AppendBytes

```go
func (path PathImpl) AppendBytes(data []byte) (err error)
```

#### func (PathImpl) Basename

```go
func (path PathImpl) Basename() string
```
Basename like path/filepath.Base

#### func (PathImpl) Canonpath

```go
func (path PathImpl) Canonpath() string
```
Canonpath retrun path with right os separator

#### func (PathImpl) Chdir

```go
func (path PathImpl) Chdir() (Path, error)
```
Chdir change current working directory do the path and return old current
working directory

#### func (PathImpl) Child

```go
func (path PathImpl) Child(childName ...string) (Path, error)
```

#### func (PathImpl) Children

```go
func (path PathImpl) Children() ([]Path, error)
```

#### func (PathImpl) CopyFile

```go
func (srcPath PathImpl) CopyFile(dst string) (p Path, err error)
```

#### func (PathImpl) CopyFrom

```go
func (path PathImpl) CopyFrom(reader io.Reader) (copyied int64, err error)
```
CopyFrom copy stream from reader to path (file after copy close and sync)

#### func (PathImpl) CryptoMd5

```go
func (path PathImpl) CryptoMd5() (hashutil.Md5, error)
```
CryptoMd5 method access hash funcionality like Path::Tiny Digest return
[hashutil.Md5](github.com/JaSei/hashutil-go) type

for example print of Md5 hexstring

    hash, err := path.CryptoMd5()
    fmt.Println(hash)

#### func (PathImpl) CryptoSha1

```go
func (path PathImpl) CryptoSha1() (hashutil.Sha1, error)
```
CryptoSha1 method access hash funcionality like Path::Tiny Digest return
[hashutil.Sha1](github.com/JaSei/hashutil-go) type

for example print of Sha1 hexstring

    hash, err := path.CryptoSha1()
    fmt.Println(hash)

#### func (PathImpl) CryptoSha256

```go
func (path PathImpl) CryptoSha256() (hashutil.Sha256, error)
```
CryptoSha256 method access hash funcionality like Path::Tiny Digest return
[hashutil.Sha256](github.com/JaSei/hashutil-go) type

for example print of Sha256 hexstring

    hash, err := path.CryptoSha256()
    fmt.Println(hash)

#### func (PathImpl) CryptoSha384

```go
func (path PathImpl) CryptoSha384() (hashutil.Sha384, error)
```
CryptoSha384 method access hash funcionality like Path::Tiny Digest return
[hashutil.Sha384](github.com/JaSei/hashutil-go) type

for example print of Sha284 hexstring

    hash, err := path.CryptoSha284()
    fmt.Println(hash)

#### func (PathImpl) CryptoSha512

```go
func (path PathImpl) CryptoSha512() (hashutil.Sha512, error)
```
CryptoSha512 method access hash funcionality like Path::Tiny Digest return
[hashutil.Sha512](github.com/JaSei/hashutil-go) type

for example print of Sha512 hexstring

    hash, err := path.CryptoSha512()
    fmt.Println(hash)

#### func (PathImpl) Exists

```go
func (path PathImpl) Exists() bool
```

#### func (PathImpl) IsDir

```go
func (path PathImpl) IsDir() bool
```
IsDir return true if path is dir

#### func (PathImpl) IsFile

```go
func (path PathImpl) IsFile() bool
```
IsFile return true is path exists and not dir (symlinks, devs, regular files)

#### func (PathImpl) IsRegularFile

```go
func (path PathImpl) IsRegularFile() bool
```
IsRegularFile return true if path is regular file (wihtout devs, symlinks, ...)

#### func (PathImpl) Lines

```go
func (path PathImpl) Lines() ([]string, error)
```
Read all lines and return as array of lines

#### func (PathImpl) LinesWalker

```go
func (path PathImpl) LinesWalker(linesFunc LinesFunc) (err error)
```
LinesWalker read lines in file and call LinesFunc with line parameter

for example:

    lines := make([]string, 0)

    linesFuncError := path.LinesWalker(func(line string) {
    	lines = append(lines, line)
    })

#### func (PathImpl) MakePath

```go
func (path PathImpl) MakePath() error
```
Make path create directory(ies) in path if not exists (like `mkdir -p`) with
default 0777 mode if you need set mode, use `MakePathMode`

#### func (PathImpl) MakePathMode

```go
func (path PathImpl) MakePathMode(mode os.FileMode) error
```
Make path create directory(ies) in path if not exists (like `mkdir -p`) with
default given mode

#### func (PathImpl) OpenReader

```go
func (path PathImpl) OpenReader() (ReadSeekCloser, error)
```
OpenReader retun ReadSeekCloser interface

for example:

    path, _ := New("/bla/bla")
    r, err := path.OpenReader()
    if err != nil {
    	panic(err)
    }
    defer r.Close()

#### func (PathImpl) OpenWriter

```go
func (path PathImpl) OpenWriter() (*os.File, error)
```
OpenWriter retun *os.File as new file (like `>>`)

for example:

    	path, _ := NewFilePath(FilePathOpt{})
    	file, err := path.OpenWriter()
    	if err != nil {
    		panic(err)
    	}
    	defer func(){
    		file.Close()
    		file.Sync()
    	}()

     writer.Write(some_bytes)

#### func (PathImpl) OpenWriterAppend

```go
func (path PathImpl) OpenWriterAppend() (*os.File, error)
```
OpenWriterAppend create new writer, similar as `OpenWriter` but append (like
`>`)

#### func (PathImpl) Parent

```go
func (path PathImpl) Parent() Path
```

    	path,_ := New("foo/bar/baz"); parent := path.Parent()   // foo/bar
     path,_ := New("foo/wible.txt"); parent := path.Parent() // foo
Returns a `Path` of corresponding to the parent directory of the original
directory or file

#### func (PathImpl) Remove

```go
func (path PathImpl) Remove() error
```
Remove file

    err := path.Remove()

like os.Remove

#### func (PathImpl) RemoveTree

```go
func (path PathImpl) RemoveTree() error
```
Remove tree of directory(ies) include files

    err := path.RemoveTree

like os.RemoveAll

#### func (PathImpl) Rename

```go
func (old PathImpl) Rename(new string) (Path, error)
```
Rename path to new path

#### func (PathImpl) Slurp

```go
func (path PathImpl) Slurp() (string, error)
```
Slurp read the whole file and return content as string

#### func (PathImpl) SlurpBytes

```go
func (path PathImpl) SlurpBytes() ([]byte, error)
```
SlurpBytes reads the whole file and return content slice of bytes like
os.ReadFile

#### func (PathImpl) Spew

```go
func (path PathImpl) Spew(content string) (err error)
```
Spew write string to file

#### func (PathImpl) SpewBytes

```go
func (path PathImpl) SpewBytes(bytes []byte) (err error)
```
SpewBytes write bytes to file

#### func (PathImpl) Stat

```go
func (path PathImpl) Stat() (os.FileInfo, error)
```
Stat return os.FileInfo

#### func (PathImpl) String

```go
func (path PathImpl) String() string
```
String return stable string representation of path this representation is linux
like (slash as separator) for os specific string use Canonpath method

#### func (PathImpl) Visit

```go
func (path PathImpl) Visit(visitFunc VisitFunc, visitOpt VisitOpt)
```

#### type ReadSeekCloser

```go
type ReadSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}
```


#### type TempOpt

```go
type TempOpt func(*tempOpt)
```

TempOpt is func for configure tempdir or tempfile

#### func  TempDir

```go
func TempDir(dir string) TempOpt
```
TempDir set directory where is temp file/dir create, empty string `""` (default)
means TEMPDIR (`os.TempDir`)

#### func  TempPrefix

```go
func TempPrefix(prefix string) TempOpt
```
TempPrefix set name beginning with prefix

#### type VisitFunc

```go
type VisitFunc func(path Path)
```


#### type VisitOpt

```go
type VisitOpt struct {
	Recurse bool
}
```


## Contributing

Contributions are very much welcome.

### Makefile

Makefile provides several handy rules, like README.md `generator` , `setup` for prepare build/dev environment, `test`, `cover`, etc...

Try `make help` for more information.

### Before pull request

please try:
* run tests (`make test`)
* run linter (`make lint`)
* if your IDE don't automaticaly do `go fmt`, run `go fmt` (`make fmt`)

### README

README.md are generate from template [.godocdown.tmpl](.godocdown.tmpl) and code documentation via [godocdown](https://github.com/robertkrimen/godocdown).

Never edit README.md direct, because your change will be lost.
