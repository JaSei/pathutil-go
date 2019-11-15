package pathutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"

	"github.com/JaSei/hashutil-go"
)

// CryptoMd5 method access hash funcionality like Path::Tiny Digest
// return [hashutil.Md5](github.com/JaSei/hashutil-go) type
//
// for example print of Md5 hexstring
//		hash, err := path.CryptoMd5()
//		fmt.Println(hash)
func (path PathImpl) CryptoMd5() (hashutil.Md5, error) {
	c, err := path.crypto(md5.New())
	if err != nil {
		return hashutil.Md5{}, nil
	}

	return hashutil.HashToMd5(c)
}

// CryptoSha1 method access hash funcionality like Path::Tiny Digest
// return [hashutil.Sha1](github.com/JaSei/hashutil-go) type
//
// for example print of Sha1 hexstring
//		hash, err := path.CryptoSha1()
//		fmt.Println(hash)
func (path PathImpl) CryptoSha1() (hashutil.Sha1, error) {
	c, err := path.crypto(sha1.New())
	if err != nil {
		return hashutil.Sha1{}, nil
	}

	return hashutil.HashToSha1(c)
}

// CryptoSha256 method access hash funcionality like Path::Tiny Digest
// return [hashutil.Sha256](github.com/JaSei/hashutil-go) type
//
// for example print of Sha256 hexstring
//		hash, err := path.CryptoSha256()
//		fmt.Println(hash)
func (path PathImpl) CryptoSha256() (hashutil.Sha256, error) {
	c, err := path.crypto(sha256.New())
	if err != nil {
		return hashutil.Sha256{}, nil
	}

	return hashutil.HashToSha256(c)
}

// CryptoSha384 method access hash funcionality like Path::Tiny Digest
// return [hashutil.Sha384](github.com/JaSei/hashutil-go) type
//
// for example print of Sha284 hexstring
//		hash, err := path.CryptoSha284()
//		fmt.Println(hash)
func (path PathImpl) CryptoSha384() (hashutil.Sha384, error) {
	c, err := path.crypto(sha512.New384())
	if err != nil {
		return hashutil.Sha384{}, nil
	}

	return hashutil.HashToSha384(c)
}

// CryptoSha512 method access hash funcionality like Path::Tiny Digest
// return [hashutil.Sha512](github.com/JaSei/hashutil-go) type
//
// for example print of Sha512 hexstring
//		hash, err := path.CryptoSha512()
//		fmt.Println(hash)
func (path PathImpl) CryptoSha512() (hashutil.Sha512, error) {
	c, err := path.crypto(sha512.New())
	if err != nil {
		return hashutil.Sha512{}, nil
	}

	return hashutil.HashToSha512(c)
}

func (path PathImpl) crypto(h hash.Hash) (ret hash.Hash, err error) {
	reader, err := path.OpenReader()
	if err != nil {
		return nil, err
	}
	defer func() {
		if errClose := reader.Close(); errClose != nil {
			err = errClose
		}
	}()

	_, err = io.Copy(h, reader)

	if err != nil {
		return nil, err
	}

	return h, nil
}
