package main

import "fmt"

func merge(firstSet, secondSet []int) (merged []int) {
	var firstPtr, secondPtr int

	for firstPtr < len(firstSet) && secondPtr < len(secondSet) {
		firstVal, secondVal := firstSet[firstPtr], secondSet[secondPtr]

		if firstVal < secondVal || secondPtr >= len(secondSet) {
			merged = append(merged, firstVal)
			firstPtr++
		} else {
			merged = append(merged, secondVal)
			secondPtr++
		}
	}

	if secondPtr < len(secondSet) {
		merged = append(merged, secondSet[secondPtr:]...)
	}

	return merged
}

func SortNumbers(numbers []int) []int {
	switch len(numbers) {
	case 0, 1:
		return numbers
	case 2:
		switch numbers[0] < numbers[1] {
		case true:
			return numbers
		default:
			return []int{numbers[1], numbers[0]}
		}
	default:
		midpt := len(numbers) / 2
		return merge(SortNumbers(numbers[:midpt]), SortNumbers(numbers[midpt:]))
	}
}

func main() {
	// fmt.Println(merge([]int{2, 3}, []int{7}))
	// fmt.Println(merge([]int{1, 3, 6, 8}, []int{2, 3, 7}))
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 5}))
	fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5}))
}

/*

merge sort base case: the array is either one or two elements long
in that case, you can easily make a comparison.

how do you merge tho?

e.g. [3, 6, 1, 8, 2, 3, 7]
split [3, 6, 1, 8] and [2, 3, 7]
-->
split [3, 6] and [1, 8]
split [2, 3] and [7]

how do you merge [2, 3] and [7]?
you compare the first indices, choose the lesser, then add that to a master array

// merge
// initialize a results array
// initialize two pointer variables to corresponsd to each array passed in
// compare the values in each array at the pointers
// 	- while the first pointer is in bounds of the first array,
// 		- if the first pointer value is less than the second pointer value,
// 		or the second pointer value does not exist
// 			- add the first pointer value to the results array
// 			- advance the first pointer by one
// 		- otherwise,
// 			- add the second pointer value to the results array
// 			- advance the second pointer by one

// return the results array

mergeSort
base case: if arr is length 1, return arr
base case: if arr is length 2,
	- if the first element is larger than the second, return [second, first]
	- otherwise, return [first, second]

otherwise,
	- return merge(first half of the arr, second half of the arr)
*/
