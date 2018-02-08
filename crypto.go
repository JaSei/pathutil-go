package pathutil

import (
	"crypto"
	"fmt"
	"hash"
	"io"
)

// Crypto method access hash funcionality like Path::Tiny Digest
// look to https://godoc.org/crypto#Hash for list possible crypto hash functions
//
// for example print of Sha256 hexstring
//		hash, err := path.Crypto(crypto.SHA256)
//		fmt.Println(hash.HexSum())

func (path PathImpl) Crypto(hash crypto.Hash) (*CryptoHash, error) {
	reader, err := path.OpenReader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	h := hash.New()

	_, err = io.Copy(h, reader)

	if err != nil {
		return nil, err
	}

	return &CryptoHash{h}, nil
}

// CryptoHash struct is only abstract for hash.Hash interface
// for possible use with methods

type CryptoHash struct {
	hash.Hash
}

// BinSum method is like hash.Sum(nil)
func (hash *CryptoHash) BinSum() []byte {
	return hash.Sum(nil)
}

// HexSum method retun hexstring representation of hash.Sum
func (hash *CryptoHash) HexSum() string {
	return fmt.Sprintf("%x", hash.Sum(nil))
}
