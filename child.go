package pathutil

import (
	"io/ioutil"
)

func (path PathImpl) Child(childName ...string) (Path, error) {
	p := make([]interface{}, len(childName)+1)
	p[0] = path
	for i, c := range childName {
		p[i+1] = c
	}

	return New(p...)
}

func (path PathImpl) Children() ([]Path, error) {
	files, err := ioutil.ReadDir(path.Canonpath())
	if err != nil {
		return nil, err
	}

	paths := make([]Path, len(files))
	for i, fileInfo := range files {
		path, _ := New(path.Canonpath(), fileInfo.Name())
		paths[i] = path
	}

	return paths, nil
}
