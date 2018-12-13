package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	temp1, err := NewTempFile(TempOpt{})
	defer func() {
		assert.NoError(t, temp1.Remove())
	}()
	assert.NotNil(t, temp1)
	assert.Nil(t, err)

	temp2, err := NewTempFile(TempOpt{Dir: "."})
	defer func() {
		assert.NoError(t, temp2.Remove())
	}()
	assert.NotNil(t, temp2)
	assert.Nil(t, err)
}

func TestTempFile(t *testing.T) {
	temp, err := NewTempFile(TempOpt{Prefix: "bla"})
	defer func() {
		assert.NoError(t, temp.Remove())
	}()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}

// Cwd is tested in chdir_test.go

func TestHome(t *testing.T) {
	home, err := Home()
	assert.NotNil(t, home)
	assert.NoError(t, err)
}
