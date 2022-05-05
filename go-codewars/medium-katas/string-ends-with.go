package main

import (
	"fmt"
)

func solution(str, ending string) bool {
	if len(str) == 0 && len(ending) != 0 {
		return false
	}

	inputEnding := str[len(str)-len(ending):]
	return inputEnding == ending
}

func main() {
	fmt.Println(solution("kittens", "ittens"))
	fmt.Println(solution("", ""))
	fmt.Println(solution(" ", ""))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("banana", "ana"))
	fmt.Println(solution("a", "z"))
	fmt.Println(solution("", "t"))
	fmt.Println(solution("$a = $b + 1", "+1"))
	fmt.Println(solution("    ", "   "))
	fmt.Println(solution("1oo", "100"))
}
