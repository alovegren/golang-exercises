package main

import (
	"fmt"
	"strings"
)

/*
Given a string, disregarding case, return a new string wherein all duplicate characters are replaced with the character ")" and all non-duplicate characters are replaced with the character "(".

Duplicates: )
Non-duplicates: (

We could first iterate through the string, making a map of the characters in the word as we go. The key in the map would be the character, while its value would be ( if it were found to be a duplicate (we can take advantage of overwriting map values) and ) otherwise. Then all that's needed in the main method is a simple lookup whilst building a new string.
*/

func buildCharacterMap(word string) map[string]string {
	characterMap := make(map[string]string)

	for _, rune := range word {
		character := string(rune)

		if _, ok := characterMap[character]; ok {
			characterMap[character] = ")" // it's a duplicate
		} else {
			characterMap[character] = "("
		}
	}

	return characterMap
}

func DuplicateEncode(word string) (encoded string) {
	lowerCaseWord := strings.ToLower(word)
	characterMap := buildCharacterMap(lowerCaseWord)

	for _, rune := range lowerCaseWord {
		character := string(rune)
		encoded += characterMap[character]
	}

	return
}

func main() {
	fmt.Println(DuplicateEncode("din"))      "((("
	fmt.Println(DuplicateEncode("recede"))   "()()()"
	fmt.Println(DuplicateEncode("Success"))  ")())())"
	fmt.Println(DuplicateEncode("(( @"))     "))(("
}
