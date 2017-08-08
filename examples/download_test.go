package pathutil_test

import (
	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	response, err := http.Get("http://example.com")

	assert.Nil(t, err)
	defer response.Body.Close()

	temp, err := pathutil.NewTempFile(pathutil.TempFileOpt{})

	defer temp.Remove()

	assert.Nil(t, err)

	copyied, err := temp.CopyFrom(response.Body)

	if !assert.Nil(t, err) {
		t.Log(err)
	}

	assert.Equal(t, int64(1270), copyied, "Copy 1270 bytes")

	ctx, err := temp.Slurp()

	assert.Nil(t, err)

	assert.Equal(t, 1270, len(ctx), "Size of http://example.com are 1270")
}
