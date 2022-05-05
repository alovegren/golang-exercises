package main

import (
	"strings"
)

func NbDig(n int, d int) (count int) {
	digitStr := fmt.Sprintf("%d", d)

	for current := 0; current <= n; current++ {
		currentStr := fmt.Sprintf("%d", current*current)
		count += strings.Count(currentStr, digitStr)
	}

	return count
}

func main() {
	fmt.Println(NbDig(10, 1))
	fmt.Println(NbDig(25, 1))
}

/*
Given an integer n and a digit d,

First, find the set of all squares from 0 to n.

Then return the number of times that the digit d appears in the set of squares.

Algorithm
Initialize a count to 0
Convert the digit d to a string
Iterate from 0 to n
	Find the square of the current value in the iteration
	Convert that square to a string
	If the resulting string contains the string version of the digit d,
		Increment the count by 1

Return the count
*/
