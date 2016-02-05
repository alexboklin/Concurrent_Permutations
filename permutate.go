package main

import (
	"fmt"
	. "go-permutations/helpers"
	"os"
	"strconv"
	//	"sync"
)

func main() {
	input := os.Args[1:]
	nums := make([]int, 0, len(input))
	for _, num := range input {
		convertedNum, _ := strconv.Atoi(num)
		nums = append(nums, convertedNum)
	}

	result := Permute(nums)

	fmt.Println("Total number of permutations: ", len(result))
	fmt.Println("Permutations: ", result)

	//	var wg sync.WaitGroup
	//
	//	results := make(chan []int, 4)
	//
	//	wg.Add(len(nums))
	//
	//	for i, currentHead := range nums {
	//
	//		go func(head int, headIndex int, input []int, output chan<- []int) {
	//			defer wg.Done()
	//
	//			tmp := make([]int, len(input))
	//			copy(tmp, input)
	//
	//			sliceWithout := append(tmp[:headIndex], tmp[headIndex+1:]...)
	//			newSlice := append([]int{head}, sliceWithout...)
	//
	//			for _, permElement := range PermTail(newSlice) {
	//				output <- permElement
	//			}
	//		}(currentHead, i, nums, results)
	//	}
	//
	//	go func() {
	//		for elem := range results {
	//			fmt.Println(elem)
	//		}
	//	}()
	//
	//	wg.Wait()
}
