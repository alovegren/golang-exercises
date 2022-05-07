package main

import (
	"fmt"
	"math"
	"sort"
)

func Comp(array1 []int, array2 []int) (areSame bool) {
	sort.Ints(array1)
	sort.Ints(array2)

	if len(array1) != len(array2) {
		return
	}

	for idx, firstElement := range array1 {
		if math.Pow(float64(firstElement), 2) != float64(array2[idx]) {
			return
		}
	}

	return true
}

func main() {
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}))
}
