package main

import (
	"fmt"
	"strings"
)

func reverseWord(word string) string {
	wordSlice := strings.Split(word, "")
	head, tail := 0, len(word)-1

	for head < tail {
		wordSlice[head], wordSlice[tail] = wordSlice[tail], wordSlice[head]

		head++
		tail--
	}

	return strings.Join(wordSlice, "")
}

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	spunWords := make([]string, len(words))

	for idx, word := range words {
		if len(word) >= 5 {
			spunWords[idx] = reverseWord(word)
		} else {
			spunWords[idx] = word
		}
	}

	return strings.Join(spunWords, " ")
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
}
