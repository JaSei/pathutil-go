/*

Pathutil is I/O utility primary inspired by David Golden's [Path::Tiny](https://metacpan.org/pod/Path::Tiny).
It is friendlier to use than [path](https://golang.org/pkg/path/), [filepath](https://golang.org/pkg/path/filepath/)
and provides many of other functions which isn't in core libraries (like `Copy` for example)

SYNOPSIS
	import (
		"fmt"
		"github.com/JaSei/pathutil-go"
	)

	// creating pathutil objects
	dir, _ := pathutil.New("/tmp")
	foo, _ := pathutil.New("foo.txt")

	subdir, _ := dir.Child("foo")
	bar, _ := subdir.Child("bar.txt")

	// stringifies as cleaned up path
	file, _ := pathutil.New("./foo.txt")
	fmt.Println(file) // "foo.txt"

	// reading files
	guts, _ := file.Slurp()
	lines, _ := file.Lines()

	// writing files
	bar.Spew(data)

	// reading directories
	children, _ := dir.Children()
	for _, child := range children {
	}

SEE ALSO

* [Path::Tiny](https://metacpan.org/pod/Path::Tiny) for Perl

* [better files](https://github.com/pathikrit/better-files) for Scala

* [pathlib](https://docs.python.org/3/library/pathlib.html) for python

*/
package pathutil
