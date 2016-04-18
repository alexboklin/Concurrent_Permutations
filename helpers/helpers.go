package helpers

/*
RotateList constructs "rotations" of a list. For instance,
RotateList [1,2,3,4] gives [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]]
*/
func RotateList(input []int) [][]int {
	return rotateListHelper(input, [][]int{}, 0)
}

func rotateListHelper(src []int, dst [][]int, acc int) [][]int {
	if acc == len(src) {
		return dst
	}
	return rotateListHelper(append(src[1:len(src)], []int{src[0]}...),
		append(dst, [][]int{src}...), acc+1)
}

/*
RotateTailAfterElem [1,2,3,4] 0: [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]], like RotateList
RotateTailAfterElem [1,2,3,4] 1: [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
RotateTailAfterElem [1,2,3,4] 2: [[1,2,3,4],[1,2,4,3]]
*/
func RotateTailAfterElem(input []int, rotateAfter int) [][]int {
	result := make([][]int, 0, len(input[rotateAfter:]))

	for _, tailRotation := range RotateList(input[rotateAfter:]) {
		temp := make([]int, len(input[:rotateAfter]))
		copy(temp, input[:rotateAfter])
		temp = append(temp, tailRotation...)
		result = append(result, temp)
	}
	return result
}

/*
Haskell-syntax formula: map $ layer n + 1 $ layer n element from rotate input
Condition: length input - 2 > 1, only one rotation for the last element
*/
func Permute(src []int) [][]int {
	currentResult := RotateTailAfterElem(src, 0)
	newResult := [][]int{}
	for i := 1; len(src)-i > 1; i++ {
		newResult = [][]int{}
		for _, element := range currentResult {
			newResult = append(newResult, RotateTailAfterElem(element, i)...)

		}
		currentResult = newResult
	}

	return currentResult
}
