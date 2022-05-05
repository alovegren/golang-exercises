package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numberStrings := strings.Split(in, " ")
	numbers := []int{}

	for _, char := range numberStrings {
		digit, _ := strconv.Atoi(char)
		numbers = append(numbers, digit)
	}

	sort.Ints(numbers)

	return fmt.Sprintf("%d %d", numbers[len(numbers)-1], numbers[0])
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))  // return "5 1"
	fmt.Println(HighAndLow("1 2 -3 4 5")) // return "5 -3"
	fmt.Println(HighAndLow("1 9 3 4 -5")) // return "9 -5"
}
