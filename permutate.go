package main

import (
	"fmt"
)

/*
rotate constructs "rotations" of a list. For instance,
rotate [1,2,3,4] gives [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]]
*/
func rotate(input []int) [][]int {
	return rotateHelper(input, [][]int{}, 0)
}

func rotateHelper(src []int, dest [][]int, accumulator int) [][]int {
	if accumulator == len(src) {
		return dest
	}
	return rotateHelper(append(src[1:len(src)], []int{src[0]}...),
		append(dest, [][]int{src}...), accumulator+1)
}

/*
permTail gives permutations of the list's tail. For instance,
permTail [1,2,3,4] gives [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
*/
func permTail(input []int) [][]int {
	result := make([][]int, 0, len(input[1:]))
	for _, tailRrotation := range rotate(input[1:]){
		result = append(result, append([]int{input[0]}, tailRrotation...))

	}
	return result
}

func main() {
	fmt.Println("rotate [1,2,3,4]: ", rotate([]int{1, 2, 3, 4}))
	fmt.Println("permTail [1,2,3,4]: ", permTail([]int{1, 2, 3, 4}))
}
