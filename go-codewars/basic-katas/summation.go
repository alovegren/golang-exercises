package main

import "fmt"

func summation(num int) int {
	sum := 0

	for addend := 1; addend <= num; addend++ {
		sum += addend
	}

	return sum
}

func main() {
	fmt.Println(summation(2))
	fmt.Println(summation(8))
}
