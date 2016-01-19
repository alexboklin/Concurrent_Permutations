package main

import "fmt"

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

func main() {
	// *rotate [1,2,3,4] gives [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]]
	fmt.Println(rotate([]int{1, 2, 3, 4}))
}
