package helpers

//import "fmt"

/*
rotate constructs "rotations" of a list. For instance,
rotate [1,2,3,4] gives [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]]
*/
func Rotate(input []int) [][]int {
	return RotateHelper(input, [][]int{}, 0)
}

func RotateHelper(src []int, dst [][]int, acc int) [][]int {
	if acc == len(src) {
		return dst
	}
	return RotateHelper(append(src[1:len(src)], []int{src[0]}...),
		append(dst, [][]int{src}...), acc+1)
}

/*
permTail gives permutations of the list's tail. For instance,
permTail [1,2,3,4] gives [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
*/
func PermTail(input []int) [][]int {
	result := make([][]int, 0, len(input[1:]))
	for _, tailRotation := range Rotate(input[1:]) {
		result = append(result, append([]int{input[0]}, tailRotation...))

	}
	return result
}

/*-- layer 0 [1,2,3,4]: [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]], like rotate
-- layer 1 [1,2,3,4]: [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
-- layer 2 [1,2,3,4]: [[1,2,3,4],[1,2,4,3]]
layer n lst = [  take n lst ++ x | x <- rotate $ drop n lst ]*/
func Layer(input []int, rotateAfter int) [][]int {
	result := make([][]int, 0, len(input[rotateAfter:]))

	for _, tailRotation := range Rotate(input[rotateAfter:]) {
		temp := make([]int, len(input[:rotateAfter]))
		copy(temp, input[:rotateAfter])
		temp = append(temp, tailRotation...)
		result = append(result, temp)
	}
	return result
}

/*-- Formula: map $ layer n + 1 $ layer n element from rotate input
-- Condition: length input - 2 > 1, only one rotation for the last element

layerize lst = layerize' (layer 0 lst) 1 where
	layerize' src counter
		| length lst - counter == 1 = src
		| otherwise = layerize' (concat $ map (layer counter) $ src) (counter + 1)*/

func Layerize(src []int) [][]int {
	currentResult := Layer(src, 0)
	newResult := [][]int{}
	for i := 1; len(src)-i > 1; i++ {
		newResult = [][]int{}
		for _, element := range currentResult {
			// TODO: run Layer() for each element concurrently
			newResult = append(newResult, Layer(element, i)...)

		}
		currentResult = newResult
	}

	//	fmt.Println(currentResult)
	return currentResult
}
