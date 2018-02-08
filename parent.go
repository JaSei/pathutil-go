package pathutil

//	path,_ := New("foo/bar/baz"); parent := path.Parent()   // foo/bar
//  path,_ := New("foo/wible.txt"); parent := path.Parent() // foo
//
// Returns a `Path` of corresponding to the parent directory of the
// original directory or file
func (path PathImpl) Parent() Path {
	// path.String() can't be empty
	parent, _ := New(path.String(), "..")
	return parent
}
