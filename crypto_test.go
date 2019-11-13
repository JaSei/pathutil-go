package pathutil

import (
	"testing"

	"github.com/JaSei/hashutil-go"
	"github.com/stretchr/testify/assert"
)

func TestPathCrypto(t *testing.T) {
	path, err := NewTempFile(TempOpt{})
	assert.Nil(t, err)
	defer func() {
		assert.NoError(t, path.Remove())
	}()

	md5Hash, err := path.CryptoMd5()
	assert.NoError(t, err)
	assert.True(t, hashutil.EmptyMd5().Equal(md5Hash))

	sha1Hash, err := path.CryptoSha1()
	assert.NoError(t, err)
	assert.True(t, hashutil.EmptySha1().Equal(sha1Hash))

	sha256Hash, err := path.CryptoSha256()
	assert.NoError(t, err)
	assert.True(t, hashutil.EmptySha256().Equal(sha256Hash))

	sha384Hash, err := path.CryptoSha384()
	assert.NoError(t, err)
	assert.True(t, hashutil.EmptySha384().Equal(sha384Hash))

	sha512Hash, err := path.CryptoSha512()
	assert.NoError(t, err)
	assert.True(t, hashutil.EmptySha512().Equal(sha512Hash))
}
