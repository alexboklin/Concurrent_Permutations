package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateList(t *testing.T) {
	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 4, 1},
		[]int{3, 4, 1, 2},
		[]int{4, 1, 2, 3},
	}, RotateList([]int{1, 2, 3, 4}))
}

func TestRotateTailAfterElem(t *testing.T) {
	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 4, 1},
		[]int{3, 4, 1, 2},
		[]int{4, 1, 2, 3},
	}, RotateTailAfterElem([]int{1, 2, 3, 4}, 0))

	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 4, 2},
		[]int{1, 4, 2, 3},
	}, RotateTailAfterElem([]int{1, 2, 3, 4}, 1))

	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 4, 3},
	}, RotateTailAfterElem([]int{1, 2, 3, 4}, 2))

}

func TestPermute(t *testing.T) {
	assert.EqualValues(t, 24, len(Permute([]int{1, 2, 3, 4})))
	assert.EqualValues(t, 120, len(Permute([]int{1, 2, 3, 4, 5})))
	assert.EqualValues(t, 720, len(Permute([]int{1, 2, 3, 4, 5, 6})))
}
