package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermTail(t *testing.T) {
	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 4, 2},
		[]int{1, 4, 2, 3},
	}, PermTail([]int{1, 2, 3, 4}))
}

func TestRotate(t *testing.T) {
	assert.EqualValues(t, [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 4, 1},
		[]int{3, 4, 1, 2},
		[]int{4, 1, 2, 3},
	}, Rotate([]int{1, 2, 3, 4}))
}

/*-- layer 0 [1,2,3,4]: [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]], like rotate
-- layer 1 [1,2,3,4]: [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
-- layer 2 [1,2,3,4]: [[1,2,3,4],[1,2,4,3]]*/
//func TestLayer(t *testing.T) {
//	assert.EqualValues(t, [][]int{
//		[]int{1, 2, 3, 4},
//		[]int{2, 3, 4, 1},
//		[]int{3, 4, 1, 2},
//		[]int{4, 1, 2, 3},
//	}, Layer([]int{1, 2, 3, 4}, 0))
//}
