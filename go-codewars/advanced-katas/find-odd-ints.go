package main

import "fmt"

func FindOdd(seq []int) (oddCount int) {
	counts := make(map[int]int)

	for _, num := range seq {
		if _, ok := counts[num]; ok {
			counts[num] += 1
		} else {
			counts[num] = 1
		}
	}

	for number, count := range counts {
		if count%2 != 0 {
			oddCount = number
		}
	}

	return
}

func main() {
	fmt.Println(FindOdd([]int{7}))
	fmt.Println(FindOdd([]int{0}))
	fmt.Println(FindOdd([]int{1, 1, 2}))
	fmt.Println(FindOdd([]int{0, 1, 0, 1, 0}))
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}

/*
iterate through the list of numbers
	if the current number has been encountered,
		add one to its count in a map
	otherwise,
		initialize it to a zero value in the map

iterate through the keys in the map
	return the first one whose value is an odd number
*/
