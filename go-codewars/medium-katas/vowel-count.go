package main

import (
	"strings"
)

func countVowels(str string) (count int) {
	chars := strings.Split(str, "")
	vowels := "aeiou"

	for _, char := range chars {
		if strings.Contains(vowels, char) {
			count += 1
		}
	}

	return count
}
