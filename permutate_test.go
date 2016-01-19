package Concurrent_Permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermTail(t *testing.T) {
	assert.EqualValues(t, permTail([]int{1, 2, 3, 4}), [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 4, 2},
		[]int{1, 4, 2, 3},
	})
}

func TestRotate(t *testing.T) {
	assert.EqualValues(t, rotate([]int{1, 2, 3, 4}), [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 4, 1},
		[]int{3, 4, 1, 2},
		[]int{4, 1, 2, 3},
	})
}
