package main

import (
	"fmt"
	"math"
)

func DigPow(n, p int) int {
	qtyDigits := len(fmt.Sprint(n))
	digits := make([]int, qtyDigits)
	whittle := n

	for position := len(digits) - 1; position >= 0; position-- {
		digits[position] = whittle % 10
		whittle /= 10
	}

	var sum float64
	for powerAdjustment, digit := range digits {
		sum += math.Pow(float64(digit), float64(p+powerAdjustment))
	}

	if int(sum)%n == 0 {
		return int(sum) / n
	} else {
		return -1
	}
}

func main() {
	fmt.Println(DigPow(89, 1))    // 1 since 8¹ + 9² = 89 = 89 * 1
	fmt.Println(DigPow(92, 1))    // -1 since there is no k such as 9¹ + 2² equals 92 * k
	fmt.Println(DigPow(695, 2))   // 2 since 6² + 9³ + 5⁴= 1390 = 695 * 2
	fmt.Println(DigPow(46288, 3)) // 51 since 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51
}
