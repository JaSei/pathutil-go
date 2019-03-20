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

func TestTempFileWithPattern(t *testing.T) {
	temp, err := NewTempFile(TempOpt{Prefix: "bla*.dat"})
	defer func() {
		assert.NoError(t, temp.Remove())
	}()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}

func TestCwd(t *testing.T) {
	cwd, err := Cwd()
	assert.NotNil(t, cwd)
	assert.NoError(t, err)

	cwdSub, err := Cwd(".git", "config")
	assert.NotNil(t, cwdSub)
	assert.NoError(t, err)

	expectedPath, _ := New(cwd.String(), ".git/config")
	assert.Equal(t, cwdSub, expectedPath)
}

func TestHome(t *testing.T) {
	home, err := Home()
	assert.NotNil(t, home)
	assert.NoError(t, err)

	homeSub, err := Home(".config", "nvim", "init.vim")
	assert.NotNil(t, homeSub)
	assert.NoError(t, err)

	expectedPath, _ := New(home.String(), ".config/nvim/init.vim")
	assert.Equal(t, homeSub, expectedPath)
}
