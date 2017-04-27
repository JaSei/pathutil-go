package path

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

func (path *Path) OpenReader() (io.Reader, *os.File, error) {
	file, err := os.Open(path.String())
	if err != nil {
		return nil, nil, err
	}

	path.file = file

	return bufio.NewReader(file), file, nil
}

func (path *Path) Slurp() ([]byte, error) {
	return ioutil.ReadFile(path.String())
}
