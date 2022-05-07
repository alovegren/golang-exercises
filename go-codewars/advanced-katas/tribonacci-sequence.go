package main

import (
	"fmt"
)

/*
Problem
Given a 'tribonacci' sequence signature i.e. a list of three numbers specifying the starting numbers in the sequence, and some positive integer n, return the first n numbers (including the signature) of the sequence.

We can solve this recursively.

The base cases can be determined based on the signature.

Though, typically fib functions return the nth fibonacci number. In this problem we're asked to collect the first n numbers.

We'll iterate while the length of the result array is less than what is needed (n).
	If the last index is less than 2, we populate it with the corresponding value in the signature array
	Otherwise, we populate it with the sum of the last three numbers in the results slice
*/

func sumNumbers(numbers []float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func Tribonacci(signature [3]float64, n int) []float64 {
	firstNElements := make([]float64, n)

	for idx := 0; idx < n; idx++ {
		if idx <= 2 {
			firstNElements[idx] = signature[idx]
		} else {
			lastThreeNumbers := firstNElements[idx-3 : idx]
			firstNElements[idx] = sumNumbers(lastThreeNumbers)
		}
	}

	return firstNElements
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))

}

// a window of the firstNElements slice, whose last index is one before the current index [idx - 2:idx + 1] and whose first index is three prior
