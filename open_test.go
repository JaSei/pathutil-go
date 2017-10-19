package pathutil

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenReader(t *testing.T) {
	path, err := NewPath("open.go")

	assert.Nil(t, err)

	reader, file, err := path.OpenReader()
	assert.Nil(t, err)
	defer file.Close()

	assert.IsType(t, new(bufio.Reader), reader)
}

func TestSlurp(t *testing.T) {
	path, err := NewPath("./LICENSE")

	assert.Nil(t, err)

	ctx, err := path.SlurpBytes()

	assert.Equal(t, 1066, len(ctx), "read LICENSE file")
}

func TestLinesFunc(t *testing.T) {
	path, err := NewPath("./LICENSE")

	assert.Nil(t, err)

	countOfLines := 0
	linesFuncError := path.LinesFunc(func(line string) {
		countOfLines++
	})

	assert.Nil(t, linesFuncError)
	assert.Equal(t, 21, countOfLines)
}

func TestLines(t *testing.T) {
	path, err := NewPath("./LICENSE")

	assert.Nil(t, err)

	lines, linesFuncError := path.Lines()

	assert.Nil(t, linesFuncError)
	assert.Equal(t, 21, len(lines))
	assert.Equal(t, "MIT License", lines[0], "string without new line on end")
}

func TestSpew(t *testing.T) {
	temp, err := NewTempFile(TempFileOpt{})
	assert.NoError(t, err)

	err = temp.Spew("kukuc")
	assert.NoError(t, err)

	ctx, err := temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "kukuc", ctx)
}
