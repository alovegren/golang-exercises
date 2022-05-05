package main

import "fmt"

func Divisors(n int) int {
	divisors := []int{}

	for candidate := 1; candidate <= n/2; candidate++ {
		if n%candidate == 0 {
			divisors = append(divisors, candidate)
		}
	}

	return len(divisors) + 1
}

func main() {
	fmt.Println(Divisors(4))
	fmt.Println(Divisors(5))
	fmt.Println(Divisors(12))
	fmt.Println(Divisors(30))
}

/*
Problem: Given a number, return the number of positive divisors that exist for that number.

e.g. the divisors of 12 are 1, 2, 3, 4, 6 and 12

Initialize a slice of integers
Iterate from 1 to halfway to the number
	If the given number, divided by the current number, results in a remainder of 0,
		Add the current number to the slice

Return the length of the slice plus one (to account for the number itself)
*/
