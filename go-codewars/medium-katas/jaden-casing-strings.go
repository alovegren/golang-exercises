package main

import (
	"fmt"
	"strings"
)

func jadenCasing(str string) string {
	words := strings.Fields(str)
	wordSlice := make([]string, len(words))

	for idx, word := range words {
		chars := strings.Split(word, "")
		capitalized := strings.ToUpper(chars[0]) + word[1:]

		wordSlice[idx] = capitalized
	}

	return strings.Join(wordSlice, " ")
}

func main() {
	fmt.Println(jadenCasing("How can mirrors be real if our eyes aren't real"))
}
