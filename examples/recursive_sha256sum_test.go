package utilpath_test

import (
	"crypto"
	"fmt"
	"github.com/JaSei/utilpath"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitRecursiveAndHashAllFiles(t *testing.T) {
	path, err := utilpath.NewPath("/tmp")
	assert.Nil(t, err)

	path.Visit(
		func(path *utilpath.Path) {
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
		utilpath.VisitOpt{Recurse: true},
	)
}
