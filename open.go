package path

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func (path *Path) OpenReader() (io.Reader, error) {
	file, err := os.Open(path.String())
	if err != nil {
		return nil, err
	}

	path.file = file

	return bufio.NewReader(file), nil
}

func (path *Path) Close() error {
	if path.file == nil {
		return nil
	}

	return path.file.Close()
}

func (path *Path) Slurp() ([]byte, error) {
	return ioutil.ReadFile(path.String())
}
