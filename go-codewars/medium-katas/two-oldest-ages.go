package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	var oldest [2]int

	sort.Ints(ages)
	copy(oldest[:], ages[len(ages)-2:])

	return oldest
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}
