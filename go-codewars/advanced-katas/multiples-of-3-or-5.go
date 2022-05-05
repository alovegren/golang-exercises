package main

import "fmt"

func Multiple3And5(number int) (sum int) {
	for count := 3; count < number; count++ {
		if count%3 == 0 || count%5 == 0 {
			sum += count
		}
	}

	return sum
}

func main() {
	fmt.Println(Multiple3And5(10))  // 23
	fmt.Println(Multiple3And5(0))   // 0
	fmt.Println(Multiple3And5(-10)) // 23
}

/*
Given an integer, return the sum of all multiples of 3 or 5 that are less than that integer

*/
