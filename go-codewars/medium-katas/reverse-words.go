package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWord(word string) string {
	head := 0
	tail := len(word) - 1

	letters := strings.Split(word, "")

	for head < tail {
		letters[head], letters[tail] = letters[tail], letters[head]
		head += 1
		tail -= 1
	}

	return strings.Join(letters, "")
}

func ReverseWords(str string) string {
	regexp, _ := regexp.Compile("(\\A|\\s)?(?P<word>\\S+)(\\s|\\z)?")

	matches := regexp.FindAllStringSubmatch(str, -1)
	strippedWordIdx := 2

	for _, match := range matches {
		word := match[strippedWordIdx]
		str = strings.Replace(str, word, reverseWord(word), 1)
	}

	return str
}

func main() {
	fmt.Println(ReverseWords("The quick brown"))
	fmt.Println(ReverseWords("stressed desserts"))
	// fmt.Println(ReverseWords("The  quick brown"))
	// fmt.Println(ReverseWords("The quick brown"))
	// fmt.Println(ReverseWords("The  quick  brown"))
}
