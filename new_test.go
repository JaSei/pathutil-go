package pathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	path, err := New("")
	assert.Nil(t, path)
	assert.Error(t, err)

	path, err = New("test")
	assert.NotNil(t, path)
	assert.NoError(t, err)

	_, err = New("test", "")
	assert.Error(t, err)
}

func TestNewTempFile(t *testing.T) {
	temp, err := NewTempFile(TempOpt{})
	defer temp.Remove()
	assert.NotNil(t, temp)
	assert.Nil(t, err)

	temp, err = NewTempFile(TempOpt{Dir: "."})
	defer temp.Remove()
	assert.NotNil(t, temp)
	assert.Nil(t, err)
}

func TestTempFile(t *testing.T) {
	temp, err := NewTempFile(TempOpt{Prefix: "bla"})
	defer temp.Remove()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}
