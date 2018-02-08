package pathutil

import (
	"io/ioutil"
)

func (path pathImpl) Child(childName ...string) (Path, error) {
	pathChunks := append([]string{path.String()}, childName...)

	return New(pathChunks...)
}

func (path pathImpl) Children() ([]Path, error) {
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
