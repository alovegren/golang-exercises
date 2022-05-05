package main

import "fmt"

func smallestInteger(arr []int) int {
	smallest := arr[0]

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] < smallest {
			smallest = arr[idx]
		}
	}

	return smallest
}

func main() {
	fmt.Println(smallestInteger([]int{34, 15, 88, 2}))
	fmt.Println(smallestInteger([]int{34, -345, -1, 100}))
}
