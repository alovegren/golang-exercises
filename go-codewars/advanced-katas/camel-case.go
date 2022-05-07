package main

import (
	"fmt"
	"regexp"
	"strings"
)

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func ToCamelCase(s string) string {
	originalWords := regexp.MustCompile("[-|_]").Split(s, -1)
	camelCasedWords := []string{}

	for idx, word := range originalWords {
		if idx == 0 {
			camelCasedWords = append(camelCasedWords, word)
		} else {
			camelCasedWords = append(camelCasedWords, capitalize(word))
		}
	}

	camelCased := strings.Join(camelCasedWords, "")

	if strings.ToUpper(string(s[0])) == string(s[0]) {
		camelCased = capitalize(camelCased)
	}

	return camelCased
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}
