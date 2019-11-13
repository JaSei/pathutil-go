package pathutil_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
)

const sizeOfExampleSite = 1256

func TestDownload(t *testing.T) {
	response, err := http.Get("http://example.com")

	if !assert.NoError(t, err) {
		t.Fatal("Connectivity problem")
	}

	defer func() {
		assert.NoError(t, response.Body.Close())
	}()

	temp, err := pathutil.NewTempFile(pathutil.TempOpt{})

	defer func() {
		assert.NoError(t, temp.Remove())
	}()

	assert.Nil(t, err)

	copyied, err := temp.CopyFrom(response.Body)

	if !assert.Nil(t, err) {
		t.Log(err)
	}

	assert.Equal(t, int64(sizeOfExampleSite), copyied, fmt.Sprintf("Copy %d bytes", sizeOfExampleSite))

	ctx, err := temp.Slurp()

	assert.Nil(t, err)

	assert.Equal(t, sizeOfExampleSite, len(ctx), fmt.Sprintf("Size of http://example.com are %d", sizeOfExampleSite))
}
