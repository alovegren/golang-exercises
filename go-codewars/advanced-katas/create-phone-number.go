package main

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	numberStrings := make([]string, 10)

	for idx, number := range numbers {
		numberStrings[idx] = fmt.Sprint(number)
	}

	areaCode := fmt.Sprintf("(%v)", strings.Join(numberStrings[0:3], ""))
	prefix := strings.Join(numberStrings[3:6], "")
	lineNumber := strings.Join(numberStrings[6:], "")

	return fmt.Sprintf("%v %v-%v", areaCode, prefix, lineNumber)
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
