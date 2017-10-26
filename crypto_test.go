package pathutil

import (
	"crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

const emptyFileSha256Hex = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

var emptyFileSha256Bin = []uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}

func TestPathCrypto(t *testing.T) {
	path, err := NewTempFile(TempOpt{})
	assert.Nil(t, err)
	defer path.Remove()

	hash, err := path.Crypto(crypto.Hash(crypto.SHA256))

	assert.Equal(t, emptyFileSha256Hex, hash.HexSum())
	assert.Equal(t, emptyFileSha256Bin, hash.BinSum())
}
