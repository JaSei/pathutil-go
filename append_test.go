package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	temp, err := NewTempFile(TempOpt{})
	assert.NoError(t, err)

	assert.NoError(t, temp.Append("test\n"))

	content, err := temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "test\n", content)

	assert.NoError(t, temp.Append("test2"))

	content, err = temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "test\ntest2", content)
}
