package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	temp, err := NewTempFile(TempOpt{})
	assert.NoError(t, err)
	assert.True(t, temp.Exists())
	assert.NoError(t, temp.Remove())
	assert.False(t, temp.Exists())
}
