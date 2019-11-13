package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenReader(t *testing.T) {
	path, err := New("open.go")

	assert.Nil(t, err)

	reader, err := path.OpenReader()
	assert.Nil(t, err)
	defer func() {
		assert.NoError(t, reader.Close())
	}()
}

func TestSlurp(t *testing.T) {
	path, err := New("./LICENSE")

	assert.Nil(t, err)

	ctx, err := path.SlurpBytes()
	assert.NoError(t, err)

	assert.Equal(t, 1066, len(ctx), "read LICENSE file")
}

func TestLinesWalker(t *testing.T) {
	path, err := New("./LICENSE")

	assert.Nil(t, err)

	countOfLines := 0
	linesFuncError := path.LinesWalker(func(line string) {
		countOfLines++
	})

	assert.Nil(t, linesFuncError)
	assert.Equal(t, 21, countOfLines)
}

func TestLines(t *testing.T) {
	path, err := New("./LICENSE")

	assert.Nil(t, err)

	lines, linesFuncError := path.Lines()

	assert.Nil(t, linesFuncError)
	assert.Equal(t, 21, len(lines))
	assert.Equal(t, "MIT License", lines[0], "string without new line on end")
}

func TestSpew(t *testing.T) {
	temp, err := NewTempFile()
	assert.NoError(t, err)

	err = temp.Spew("kukuc")
	assert.NoError(t, err)

	ctx, err := temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "kukuc", ctx)
}
