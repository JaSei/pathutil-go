package pathutil

//	path,_ := NewPath("foo/bar/baz"); parent := path.Parent()   // foo/bar
//  path,_ := NewPath("foo/wible.txt"); parent := path.Parent() // foo
//
// Returns a `Path` of corresponding to the parent directory of the
// original directory or file
func (path pathImpl) Parent() Path {
	// path.String() can't be empty
	parent, _ := NewPath(path.String(), "..")
	return parent
}
