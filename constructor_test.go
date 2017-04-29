package pathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPath(t *testing.T) {
	path, err := NewPath("")
	assert.Nil(t, path)
	assert.NotNil(t, err)

	path, err = NewPath("test")
	assert.NotNil(t, path)
	assert.Nil(t, err)
}

func TestNewTempFile(t *testing.T) {
	temp, err := NewTempFile(TempFileOpt{})
	defer temp.Remove()
	assert.NotNil(t, temp)
	assert.Nil(t, err)

	temp, err = NewTempFile(TempFileOpt{Dir: "."})
	defer temp.Remove()
	assert.NotNil(t, temp)
	assert.Nil(t, err)
}

func TestTempFile(t *testing.T) {
	temp, err := NewTempFile(TempFileOpt{Prefix: "bla"})
	defer temp.Remove()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}
