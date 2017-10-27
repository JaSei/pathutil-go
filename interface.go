package pathutil

import (
	"crypto"
	"io"
	"os"
)

type VisitFunc func(path Path)

type VisitOpt struct {
	Recurse bool
}

type LinesFunc func(string)

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

	OpenReader() (io.Reader, *os.File, error)
	OpenWriter() (*os.File, error)

	Slurp() (string, error)
	SlurpBytes() ([]byte, error)

	Spew(string) error
	SpewBytes([]byte) error

	Lines() ([]string, error)
	LinesWalker(LinesFunc) error

	Child(...string) (Path, error)
	Children() ([]Path, error)
}
