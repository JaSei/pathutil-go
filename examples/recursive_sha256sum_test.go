package pathutil_test

import (
	"fmt"
	"testing"

	"github.com/JaSei/pathutil-go"
	"github.com/stretchr/testify/assert"
)

func TestVisitRecursiveAndHashAllFiles(t *testing.T) {
	path, err := pathutil.New("tree")
	assert.Nil(t, err)

	path.Visit(
		func(path pathutil.Path) {
			if path.IsDir() {
				return
			}

			hash, err := path.CryptoSha256()

			if err == nil {
				fmt.Printf("%s\t%s\n", hash, path)
			} else {
				fmt.Println(err)
			}
		},
		pathutil.VisitOpt{Recurse: true},
	)
}
