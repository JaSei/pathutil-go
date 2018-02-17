package pathutil_test

import (
	"fmt"
	"testing"

	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
)

func TestSynopsis(t *testing.T) {
	// creating pathutil objects
	dir, _ := pathutil.New("/tmp")
	//foo, _ := pathutil.New("foo.txt")

	subdir, _ := dir.Child("foo")
	assert.Equal(t, "/tmp/foo", subdir.String())
	bar, _ := subdir.Child("bar.txt")
	assert.Equal(t, "/tmp/foo/bar.txt", bar.String())

	file, _ := pathutil.New("./foo.txt")
	assert.Equal(t, "foo.txt", file.String())
	fmt.Println(file)

	// reading files
	guts, _ := file.Slurp()
	assert.Empty(t, guts)
	lines, _ := file.Lines()
	assert.Empty(t, lines)

	// writing files
	_ = bar.Spew("ahoj")

	// reading directories
	children, _ := dir.Children()
	for _, child := range children {
		fmt.Println(child)
	}
}
