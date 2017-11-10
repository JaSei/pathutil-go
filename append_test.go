package pathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	temp, err := NewTempFile(TempOpt{})
	assert.NoError(t, err)

	temp.Append("test\n")

	content, err := temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "test\n", content)

	temp.Append("test2")

	content, err = temp.Slurp()
	assert.NoError(t, err)
	assert.Equal(t, "test\ntest2", content)
}
