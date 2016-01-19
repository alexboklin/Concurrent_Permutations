package Concurrent_Permutations

import (
	"fmt"
	"sync"
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

func main() {
	var wg sync.WaitGroup

	testInput := []int{1, 2, 3, 4}

	results := make(chan []int, 4)

	wg.Add(len(testInput))

	for i, currentHead := range testInput {

		go func(head int, headIndex int, input []int, output chan<- []int){
			defer wg.Done()

			tmp := make([]int, len(input))
			copy(tmp, input)

			sliceWithout := append(tmp[:headIndex], tmp[headIndex+1:]...)
			newSlice := append([]int{head}, sliceWithout...)

			for _, permElement := range permTail(newSlice) {
				output <- permElement
			}
		}(currentHead, i, testInput, results)
	}

	go func() {
		for elem := range results {
			fmt.Println(elem)
		}
	}()

	wg.Wait()
}
