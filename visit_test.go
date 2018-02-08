package pathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitFlat(t *testing.T) {
	path, err := New("examples/tree")

	assert.Nil(t, err)
	flat_expected := map[string]int{
		"examples/tree/4": 0,
	}

	flat := make(map[string]int)
	path.Visit(
		func(path Path) {
			flat[path.String()] = 0
		},
		VisitOpt{},
	)

	assert.Equal(t, flat_expected, flat, "flat files")
}

func TestVisitRecurse(t *testing.T) {
	path, err := New("examples/tree")

	assert.Nil(t, err)

	flat_expected := map[string]int{
		"examples/tree/a":     0,
		"examples/tree/a/1":   0,
		"examples/tree/a/b":   0,
		"examples/tree/a/b/2": 0,
		"examples/tree/c":     0,
		"examples/tree/c/3":   0,
		"examples/tree/4":     0,
	}

	flat := make(map[string]int)
	path.Visit(
		func(path Path) {
			flat[path.String()] = 0
		},
		VisitOpt{Recurse: true},
	)

	assert.Equal(t, flat_expected, flat, "flat files")

}
