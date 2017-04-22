package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisit(t *testing.T) {
	path, err := NewPath("examples/tree")

	assert.Nil(t, err)

	t.Run("flat visit", func(t *testing.T) {
		flat_expected := map[string]int{
			"examples/tree/4": 0,
		}

		flat := make(map[string]int)
		path.Visit(
			func(path *Path) {
				flat[path.String()] = 0
			},
			VisitOpt{},
		)

		assert.Equal(t, flat_expected, flat, "flat files")
	})

	t.Run("flat recurse", func(t *testing.T) {
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
			func(path *Path) {
				flat[path.String()] = 0
			},
			VisitOpt{Recurse: true},
		)

		assert.Equal(t, flat_expected, flat, "flat files")
	})

}
