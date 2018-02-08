# pathutil

[![Release](https://img.shields.io/github/release/JaSei/pathutil-go.svg?style=flat-square)](https://github.com/JaSei/pathutil-go/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/JaSei/pathutil-go.svg?style=flat-square)](https://travis-ci.org/JaSei/pathutil-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/JaSei/pathutil-go?style=flat-square)](https://goreportcard.com/report/github.com/JaSei/pathutil-go)
[![GoDoc](https://godoc.org/github.com/JaSei/pathutil-go?status.svg&style=flat-square)](http://godoc.org/github.com/JaSei/pathutil-go)
[![codecov.io](https://codecov.io/github/JaSei/pathutil-go/coverage.svg?branch=master)](https://codecov.io/github/JaSei/pathutil-go?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/JaSei/pathutil-go/-/badge.svg)](https://sourcegraph.com/github.com/JaSei/pathutil-go?badge)



## Usage

#### type CryptoHash

```go
type CryptoHash struct {
	hash.Hash
}
```


#### func (*CryptoHash) BinSum

```go
func (hash *CryptoHash) BinSum() []byte
```
BinSum method is like hash.Sum(nil)

#### func (*CryptoHash) HexSum

```go
func (hash *CryptoHash) HexSum() string
```
HexSum method retun hexstring representation of hash.Sum

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

	Crypto(crypto.Hash) (*CryptoHash, error)

	MakePath() error
	MakePathMode(os.FileMode) error

	OpenReader() (io.ReadCloser, error)
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
func Cwd() (Path, error)
```
Cwd create new path from current working directory

#### func  New

```go
func New(path ...string) (Path, error)
```
New construct Path

for example

    path := New("/home/test", ".vimrc")

if you can use `Path` in `New`, you must use `.String()` method

#### func  NewTempDir

```go
func NewTempDir(options TempOpt) (Path, error)
```
NewTempDir create temp directory

for cleanup use `defer`

    	tempdir, err := pathutil.NewTempDir(TempOpt{})
     defer tempdir.RemoveTree()

#### func  NewTempFile

```go
func NewTempFile(options TempOpt) (Path, error)
```

#### type TempOpt

```go
type TempOpt struct {
	// directory where is temp file/dir create, empty string `""` (default) means TEMPDIR (`os.TempDir`)
	Dir string
	// name beginning with prefix
	Prefix string
}
```

TempOpt is struct for configure new tempdir or tempfile

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
