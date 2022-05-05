package main

import (
	"fmt"
	"math"
)

func makeNegative(x int) int {
	absValue := math.Abs(float64(x))
	return int(-1 * absValue)
}

func main() {
	fmt.Println(makeNegative(-44))
}
