package pathutil_test

import (
	"crypto"
	"fmt"
	"github.com/JaSei/pathutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitRecursiveAndHashAllFiles(t *testing.T) {
	path, err := pathutil.NewPath("/tmp")
	assert.Nil(t, err)

	path.Visit(
		func(path *pathutil.Path) {
			if path.IsDir() {
				return
			}

			hash, err := path.Crypto(crypto.SHA256)

			if err == nil {
				fmt.Printf("%s\t%s\n", hash.HexSum(), path.String())
			} else {
				fmt.Println(err)
			}
		},
		pathutil.VisitOpt{Recurse: true},
	)
}
