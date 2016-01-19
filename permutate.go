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
	for _, tailRotation := range rotate(input[1:]) {
		result = append(result, append([]int{input[0]}, tailRotation...))

	}
	return result
}

func worker(head int, headIndex int, input []int, output chan<- []int) {
	//	fmt.Println("head: ", head, "input: ", input)

	tmp := make([]int, len(input))
	copy(tmp, input)

	sliceWithout := append(tmp[:headIndex], tmp[headIndex+1:]...)
	//	fmt.Println("without: ", sliceWithout)
	newSlice := append([]int{head}, sliceWithout...)
	//	fmt.Println("new: ", newSlice)
	//	fmt.Println()

	for _, permElement := range permTail(newSlice) {
//		fmt.Println("permElement: ", permElement)
		output <- permElement
	}

	if headIndex == len(input)-1 {
		close(output)
	}

}

func main() {
	testInput := []int{1, 2, 3, 4}

	results := make(chan []int, 4)

	//	fmt.Println("rotate [1,2,3,4]: ", rotate(testInput))
	//	fmt.Println("permTail [1,2,3,4]: ", permTail(testInput))

	for i, currentHead := range testInput {
		/*		tmp := make([]int, len(testInput))
				copy(tmp, testInput)

				fmt.Println(currentHead)
				sliceWithout := append(tmp[:i], tmp[i+1:]...)
				fmt.Println("without: ", sliceWithout)
				newSlice := append([]int{currentHead}, sliceWithout...)
				fmt.Println("new: ", newSlice)

				fmt.Println()*/

		go worker(currentHead, i, testInput, results)
	}

	//	<-results

	for elem := range results {
		fmt.Println(elem)
	}

}
