package pathutil

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// OpenReader retun io.ReaderCloser
//
// for example:
//	path, _ := New("/bla/bla")
//	r, err := path.OpenReader()
//	if err != nil {
//		panic(err)
//	}
//	defer r.Close()
//
func (path pathImpl) OpenReader() (io.ReadCloser, error) {
	file, err := os.Open(path.Canonpath())
	if err != nil {
		return nil, errors.Wrapf(err, "OpenReader on path %s fail (%+v)", path, path)
	}

	return file, nil
}

// OpenWriter retun *os.File as new file (like `>>`)
//
// for example:
//	path, _ := NewFilePath(FilePathOpt{})
//	file, err := path.OpenWriter()
//	if err != nil {
//		panic(err)
//	}
//	defer func(){
//		file.Close()
//		file.Sync()
//	}()
//
//  writer.Write(some_bytes)
//
func (path pathImpl) OpenWriter() (*os.File, error) {
	return path.openWriterFlag(os.O_RDWR | os.O_CREATE)
}

// OpenWriterAppend create new writer, similar as `OpenWriter` but append (like `>`)
func (path pathImpl) OpenWriterAppend() (*os.File, error) {
	return path.openWriterFlag(os.O_APPEND | os.O_CREATE | os.O_WRONLY)
}

func (path pathImpl) openWriterFlag(flag int) (*os.File, error) {
	file, err := os.OpenFile(path.String(), flag, 0775)
	if err != nil {
		return nil, err
	}

	return file, err
}

// Slurp read all file
// like ioutil.ReadFile
func (path pathImpl) Slurp() (string, error) {
	bytes, err := path.SlurpBytes()
	if err != nil {
		return "", err
	}

	return string(bytes[:]), nil
}

func (path pathImpl) SlurpBytes() ([]byte, error) {
	return ioutil.ReadFile(path.String())
}

// Spew write string to file
func (path pathImpl) Spew(content string) error {
	file, err := path.OpenWriter()
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}

// SpewBytes write bytes to file
func (path pathImpl) SpewBytes(bytes []byte) error {
	file, err := path.OpenWriter()
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := file.Write(bytes); err != nil {
		return err
	}

	return nil
}

// LinesWalker read lines in file and call LinesFunc with line parameter
//
// for example:
//	lines := make([]string, 0)
//
//	linesFuncError := path.LinesWalker(func(line string) {
//		lines = append(lines, line)
//	})
func (path pathImpl) LinesWalker(linesFunc LinesFunc) error {
	reader, err := path.OpenReader()
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		linesFunc(scanner.Text())
	}

	return scanner.Err()
}

// Read all lines and return as array of lines
func (path pathImpl) Lines() ([]string, error) {
	lines := make([]string, 0)

	linesFuncError := path.LinesWalker(func(line string) {
		lines = append(lines, line)
	})

	return lines, linesFuncError
}

// CopyFrom copy stream from reader to path
// (file after copy close and sync)
func (path pathImpl) CopyFrom(reader io.Reader) (int64, error) {
	file, err := path.OpenWriter()
	if err != nil {
		return 0, err
	}

	defer func() {
		file.Close()
		file.Sync()
	}()

	copyied, err := io.Copy(file, reader)

	return copyied, err
}
