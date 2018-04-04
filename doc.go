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

BREAKING CHANGE 0.2.0 -> 1.0.0

`NewTempFile` or `NewTempDir` don't need TempOpt struct

	//0.2.0 default
	pathutil.NewTempFile(pathutil.TempOpt{})
	//0.2.0 custom
	pathutil.NewTempFile(pathutil.TempOpt{Dir: "/test", Prefix: "pre"})

	//1.0.0 default
	pathutil.NewTempFile()
	//1.0.0 custom
	pathutil.NewTempFile(Dir("/test"), Prefix("pre"))

`New` method parameter allowed `string` type or type implements `fmt.Stringer` interface

	//0.2.0
	pathutil.New(otherPath.String(), "test")

	//1.0.0
	pathutil.New(otherPath, "test")

This shouldn't be breaking change, but if you use in some code variadic parameter as input of `pathutil.New`, then can be problem

	//0.2.0
	func(p ...string) {
		pathutil.New(p...)
	}("a", "b")

	//1.0.0
	func(p ...string) {
		n := make([]interface{}, len(p))
		for i, v := range p {
			n[i] = v
		}
		pathutil.New(n...)
	}("a", "b")

*/
package pathutil
