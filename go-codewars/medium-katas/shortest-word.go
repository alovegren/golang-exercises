package main

import (
	"fmt"
	"strings"
)

func shortestWord(str string) int {
	words := strings.Fields(str)
	shortestLength := len(words[0])

	for _, word := range words[1:] {
		length := len(word)
		if length < shortestLength {
			shortestLength = length
		}
	}

	return shortestLength
}

func main() {
	fmt.Println(shortestWord("peter piper ate a peck")) // 1
}
