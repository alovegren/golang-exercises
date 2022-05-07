package main

import (
	"fmt"
	"math"
)

/*
Given an integer m, representing the total volume of a building, find the number of cubes n of which the building is constructed.

The volume of each cubed is determined by the following sequence, where the bottom-most cube is represented by the left-most expression:

(n^3) + ((n-1)^3) + ((n-2)^3) + ((n-3)^3) ... 1^3 = 1

In order for the last cube's volume to be 1, we have to find some p such that
(n-p) = 1

If we're given the volume, we know that the volume of the remaining part of the sequence is (volume - 1)
*/

func FindNb(m int) int {
	var sum int
	var n int

	for ; sum < m; n++ {
		sum += int(math.Pow(float64(n), 3))
	}

	if sum == m {
		return n - 1
	}

	return -1
}

func main() {
	fmt.Println(FindNb(1071225))
	fmt.Println(FindNb(91716553919377))
	fmt.Println(FindNb(4183059834009))
	fmt.Println(FindNb(24723578342962))
}
