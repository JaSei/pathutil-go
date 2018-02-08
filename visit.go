package pathutil

import (
	"os"
	"path/filepath"
)

func (path PathImpl) Visit(visitFunc VisitFunc, visitOpt VisitOpt) {
	walkFn := func(file string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() && file != path.String() && !visitOpt.Recurse {
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}

		//skip self path
		if file == path.String() {
			return nil
		}

		path, err := New(file)

		visitFunc(path)

		return nil
	}

	filepath.Walk(path.String(), walkFn)
}
