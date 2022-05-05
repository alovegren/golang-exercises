package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func squareEveryDigit(inputNum int) string {
	inputStr := strconv.Itoa(inputNum)
	squared := ""

	for _, char := range strings.Split(inputStr, "") {
		digit, _ := strconv.ParseFloat(char, 64)
		squaredChar := fmt.Sprintf("%g", math.Pow(digit, 2))
		squared += squaredChar
	}

	return squared
}

func main() {
	squareEveryDigit(315)
}
