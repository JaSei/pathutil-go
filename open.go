package pathutil

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// OpenReader retun bufio io.Reader
// because bufio io.Reader don't implement Close method,
// OpenReader returns *os.File too
//
// for example:
//	path, _ := NewPath("/bla/bla")
//	reader, file, err := path.OpenReader()
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
func (path *Path) OpenReader() (io.Reader, *os.File, error) {
	file, err := os.Open(path.String())
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewReader(file), file, nil
}

// OpenWriter retun bufio io.Writer
// because bufio io.Writer don't implement Close method,
// OpenWriter returns *os.File too
//
// for example:
//	path, _ := NewFilePath(FilePathOpt{})
//	writer, file, err := path.OpenWriter()
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
func (path *Path) OpenWriter() (io.Writer, *os.File, error) {
	file, err := os.Open(path.String())
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewWriter(file), file, nil
}

// Slurp read all file
// like ioutil.ReadFile
func (path *Path) Slurp() ([]byte, error) {
	return ioutil.ReadFile(path.String())
}

// Read lines in file and call linesFunc with line parameter
//
// for example:
//	lines := make([]string, 0)
//
//	linesFuncError := path.LinesFunc(func(line string) {
//		lines = append(lines, line)
//	})
func (path *Path) LinesFunc(linesFunc func(string)) error {
	reader, file, err := path.OpenReader()
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		linesFunc(scanner.Text())
	}

	return scanner.Err()
}

// Read all lines and return as array of lines
func (path *Path) Lines() ([]string, error) {
	lines := make([]string, 0)

	linesFuncError := path.LinesFunc(func(line string) {
		lines = append(lines, line)
	})

	return lines, linesFuncError
}

// CopyFrom copy stream from reader to path
// (file after copy close and sync)
func (path *Path) CopyFrom(reader io.Reader) (int64, error) {
	writer, file, err := path.OpenWriter()
	if err != nil {
		return 0, err
	}

	defer func() {
		file.Close()
		file.Sync()
	}()

	copyied, err := io.Copy(writer, reader)

	return copyied, err
}
