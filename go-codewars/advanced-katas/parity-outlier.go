package main

import (
	"fmt"
	"math"
)

func getParityCode(integers []int) int {
	isEven := 0

	for _, integer := range integers {
		if integer%2 == 0 {
			isEven += 1
		}
	}

	switch {
	case isEven >= 2:
		return 0
	default:
		return 1
	}
}

func FindOutlier(integers []int) int {
	parityCode := getParityCode(integers[0:3])

	for _, integer := range integers {
		if math.Abs(float64(integer%2)) != float64(parityCode) {
			return integer
		}
	}

	return 0
}

func main() {
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))         // 11
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))         // 160
	fmt.Println(FindOutlier([]int{11, 3, 1719, 11, -22, 19, 11, 13, -21})) // -22
}
