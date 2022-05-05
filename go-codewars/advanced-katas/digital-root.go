package main

import (
	"strconv"
	"strings"
)

func getDigitSum(n int) (sum int) {
	for _, rune := range strings.Split(strconv.Itoa(n), "") {
		digit, _ := strconv.Atoi(rune)
		sum += digit
	}

	return
}

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	} else {
		return DigitalRoot(getDigitSum(n))
	}
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(942))
	fmt.Println(DigitalRoot(132189))
	fmt.Println(DigitalRoot(493193))
}
