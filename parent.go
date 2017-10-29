package pathutil

func (path pathImpl) Parent() (Path, error) {
	return NewPath(path.String(), "..")
}
