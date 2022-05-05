package main

import (
	"fmt"
)

func IsTriangle(a, b, c int) bool {
	switch {
	case a <= 0, b <= 0, c <= 0:
		return false
	case (a+b <= c), (a+c <= b), (b+c <= a):
		return false
	default:
		return true
	}
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(1, 2, 5))
	fmt.Println(IsTriangle(2, 5, 1))
}
