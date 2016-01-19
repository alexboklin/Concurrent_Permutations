package main

import (
	. "concurrent_permutations/helpers"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	testInput := []int{1, 2, 3, 4}

	results := make(chan []int, 4)

	wg.Add(len(testInput))

	for i, currentHead := range testInput {

		go func(head int, headIndex int, input []int, output chan<- []int) {
			defer wg.Done()

			tmp := make([]int, len(input))
			copy(tmp, input)

			sliceWithout := append(tmp[:headIndex], tmp[headIndex+1:]...)
			newSlice := append([]int{head}, sliceWithout...)

			for _, permElement := range PermTail(newSlice) {
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
