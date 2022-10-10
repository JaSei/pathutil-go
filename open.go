package pathutil

import (
	"bufio"
	"io"
	"os"

	"github.com/pkg/errors"
)

type ReadSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

// OpenReader retun ReadSeekCloser interface
//
// for example:
//
//	path, _ := New("/bla/bla")
//	r, err := path.OpenReader()
//	if err != nil {
//		panic(err)
//	}
//	defer r.Close()
func (path PathImpl) OpenReader() (ReadSeekCloser, error) {
	file, err := os.Open(path.Canonpath())
	if err != nil {
		return nil, errors.Wrapf(err, "OpenReader on path %s fail (%+v)", path, path)
	}

	return file, nil
}

// OpenWriter retun *os.File as new file (like `>>`)
//
// for example:
//
//		path, _ := NewFilePath(FilePathOpt{})
//		file, err := path.OpenWriter()
//		if err != nil {
//			panic(err)
//		}
//		defer func(){
//			file.Close()
//			file.Sync()
//		}()
//
//	 writer.Write(some_bytes)
func (path PathImpl) OpenWriter() (*os.File, error) {
	return path.openWriterFlag(os.O_RDWR | os.O_CREATE)
}

// OpenWriterAppend create new writer, similar as `OpenWriter` but append (like `>`)
func (path PathImpl) OpenWriterAppend() (*os.File, error) {
	return path.openWriterFlag(os.O_APPEND | os.O_CREATE | os.O_WRONLY)
}

func (path PathImpl) openWriterFlag(flag int) (*os.File, error) {
	file, err := os.OpenFile(path.String(), flag, 0775)
	if err != nil {
		return nil, err
	}

	return file, err
}

// Slurp read the whole file and return content as string
func (path PathImpl) Slurp() (string, error) {
	bytes, err := path.SlurpBytes()
	if err != nil {
		return "", err
	}

	return string(bytes[:]), nil
}

// SlurpBytes reads the whole file and return content slice of bytes
// like os.ReadFile
func (path PathImpl) SlurpBytes() ([]byte, error) {
	return os.ReadFile(path.String())
}

// Spew write string to file
func (path PathImpl) Spew(content string) (err error) {
	file, err := path.OpenWriter()
	if err != nil {
		return err
	}

	defer func() {
		if errClose := file.Close(); errClose != nil {
			err = errClose
		}
	}()

	_, err = file.WriteString(content)
	return err
}

// SpewBytes write bytes to file
func (path PathImpl) SpewBytes(bytes []byte) (err error) {
	file, err := path.OpenWriter()
	if err != nil {
		return err
	}

	defer func() {
		if errClose := file.Close(); errClose != nil {
			err = errClose
		}
	}()

	_, err = file.Write(bytes)
	return err
}

// LinesWalker read lines in file and call LinesFunc with line parameter
//
// for example:
//
//	lines := make([]string, 0)
//
//	linesFuncError := path.LinesWalker(func(line string) {
//		lines = append(lines, line)
//	})
func (path PathImpl) LinesWalker(linesFunc LinesFunc) (err error) {
	reader, err := path.OpenReader()
	if err != nil {
		return err
	}
	defer func() {
		if errClose := reader.Close(); errClose != nil {
			err = errClose
		}
	}()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		linesFunc(scanner.Text())
	}

	return scanner.Err()
}

// Read all lines and return as array of lines
func (path PathImpl) Lines() ([]string, error) {
	lines := make([]string, 0)

	linesFuncError := path.LinesWalker(func(line string) {
		lines = append(lines, line)
	})

	return lines, linesFuncError
}

// CopyFrom copy stream from reader to path
// (file after copy close and sync)
func (path PathImpl) CopyFrom(reader io.Reader) (copyied int64, err error) {
	file, err := path.OpenWriter()
	if err != nil {
		return 0, err
	}

	defer func() {
		if errSync := file.Sync(); errSync != nil {
			err = errSync
		} else if errClose := file.Close(); errClose != nil {
			err = errClose
		}

	}()

	copyied, err = io.Copy(file, reader)

	return copyied, err
}
