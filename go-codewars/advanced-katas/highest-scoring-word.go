package main

import (
	"fmt"
	"strings"
)

func scoreWord(word string) (score int) {
	ALPHA_OFFSET := 96

	for _, rune := range word {
		score += (int(rune) - ALPHA_OFFSET)
	}

	return
}

func High(s string) string {
	s = strings.ToLower(s)
	words := strings.Split(s, " ")

	var highestScore int
	var highestScoringWord string

	for _, word := range words {
		score := scoreWord(word)

		if score > highestScore {
			highestScore = score
			highestScoringWord = word
		}
	}

	fmt.Println(highestScoringWord)
	return highestScoringWord
}

func main() {
	High("man i need a taxi up to ubud")
}
