package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) // the `int64` preceding parentheses is a type conversion

	fmt.Println(math.Sin(n)) // `Sin` expects a `float64` so `n` is given this type
}