package main

import (
	"fmt"
)

func Solution(str string) []string {
	pairs := []string{}

	for startIdx := 0; startIdx < len(str); startIdx += 2 {
		if startIdx+1 >= len(str) {
			pairs = append(pairs, string(str[startIdx])+"_")
		} else {
			pairs = append(pairs, str[startIdx:startIdx+2])
		}

	}

	return pairs
}

func main() {
	Solution("abc")    // ['ab', 'c_']
	Solution("abcdef") // ['ab', 'cd', 'ef']
}
