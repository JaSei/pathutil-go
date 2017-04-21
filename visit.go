package path

import (
	"os"
	"path/filepath"
)

type VisitFunc func(path *Path)

type VisitOpt struct {
	Recurse bool
}

func (path *Path) Visit(visitFunc VisitFunc, visitOpt VisitOpt) {
	walkFn := func(file string, info os.FileInfo, err error) error {
		if info != nil && info.IsDir() && file != path.String() && !visitOpt.Recurse {
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}

		path, err := NewPath(file)

		visitFunc(path)

		return nil
	}

	filepath.Walk(path.String(), walkFn)
}
