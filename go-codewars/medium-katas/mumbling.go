package main

import (
	"fmt"
	"strings"
)

func accum(str string) string {
	chars := strings.Split(strings.ToLower(str), "")
	groups := []string{}

	for lowerCount, char := range chars {
		groups = append(
			groups,
			strings.ToUpper(char)+strings.Repeat(char, lowerCount),
		)
	}

	return strings.Join(groups, "-")
}

func main() {
	fmt.Println(accum("abcd"))    // "A-Bb-Ccc-Dddd"
	fmt.Println(accum("RqaEzty")) // "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
	fmt.Println(accum("cwAt"))    // "C-Ww-Aaa-Tttt"
}
