package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) (duplicateCount int) {
	letterCounts := make(map[string]int)

	for _, rune := range s1 {
		letter := strings.ToLower(string(rune))
		letterCount := letterCounts[letter]

		if letterCount == 1 {
			duplicateCount++
		} else {
			letterCounts[letter] += 1
		}
	}

	return
}

func main() {
	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count("aabbcde"))
	fmt.Println(duplicate_count("aabbcde"))
	fmt.Println(duplicate_count("abcdeaB11"))
}
