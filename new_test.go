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

	_, err = New("test", nil)
	assert.Error(t, err)

	_, err = New("test", 64)
	assert.Error(t, err)

	func(p ...string) {
		n := make([]interface{}, len(p))
		for i, v := range p {
			n[i] = v
		}
		path, err = New(n...)
		assert.NoError(t, err)
		assert.Equal(t, path.String(), "a/b")
	}("a", "b")

}

func TestNewTempFile(t *testing.T) {
	temp1, err := NewTempFile()
	defer func() {
		assert.NoError(t, temp1.Remove())
	}()
	assert.NotNil(t, temp1)
	assert.NoError(t, err)

	temp2, err := NewTempFile(TempDir("."))
	defer func() {
		assert.NoError(t, temp2.Remove())
	}()
	assert.NotNil(t, temp2)
	assert.NoError(t, err)

	temp, err := NewTempFile(TempPrefix("bla"))
	defer func() {
		assert.NoError(t, temp.Remove())
	}()

	assert.NotNil(t, temp)
	assert.NoError(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
}

func TestCwd(t *testing.T) {
	cwd, err := Cwd()
	assert.NotNil(t, cwd)
	assert.NoError(t, err)

	cwdSub, err := Cwd(".git", "config")
	assert.NotNil(t, cwdSub)
	assert.NoError(t, err)

	expectedPath, _ := New(cwd, ".git/config")
	assert.Equal(t, cwdSub, expectedPath)
}

func TestHome(t *testing.T) {
	home, err := Home()
	assert.NotNil(t, home)
	assert.NoError(t, err)

	homeSub, err := Home(".config", "nvim", "init.vim")
	assert.NotNil(t, homeSub)
	assert.NoError(t, err)

	expectedPath, _ := New(home, ".config/nvim/init.vim")
	assert.Equal(t, homeSub, expectedPath)
}

func TestTempFileWithPattern(t *testing.T) {
	temp, err := NewTempFile(TempPrefix("bla*.dat"))
	defer func() {
		assert.NoError(t, temp.Remove())
	}()

	assert.NotNil(t, temp)
	assert.Nil(t, err)
	assert.Exactly(t, true, temp.Exists(), "new temp file exists")
	assert.Regexp(t, "bla.+\\.dat$", temp.String())
}

func TestNewTempDir(t *testing.T) {
	t.Run("new temp dir", func(t *testing.T) {
		temp, err := NewTempDir()
		defer func() {
			assert.NoError(t, temp.RemoveTree())
		}()
		assert.NotNil(t, temp)
		assert.NoError(t, err)
	})

	t.Run("new temp dir in current directory", func(t *testing.T) {
		temp2, err := NewTempDir(TempDir("."))
		defer func() {
			assert.NoError(t, temp2.RemoveTree())
		}()
		assert.NotNil(t, temp2)
		assert.NoError(t, err)
	})

	t.Run("new temp dir with prefix", func(t *testing.T) {
		temp, err := NewTempDir(TempPrefix("bla"))
		defer func() {
			assert.NoError(t, temp.RemoveTree())
		}()

		assert.NotNil(t, temp)
		assert.NoError(t, err)
		assert.Exactly(t, true, temp.Exists(), "new temp dir exists")
	})
}
