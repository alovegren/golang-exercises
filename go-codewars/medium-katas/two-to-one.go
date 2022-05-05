package main

import (
	"fmt"
	"sort"
	"strings"
)

func longest(s1, s2 string) (longest string) {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)

	uniqueChars := make(map[string]bool)
	for _, char := range chars {
		if !uniqueChars[char] {
			longest += char
		}
		uniqueChars[char] = true
	}

	return longest
}

func main() {
	fmt.Println(longest("xyaabbbccccdefww", "xxxxyyyyabklmopq") == "abcdefklmopqwxy")
	fmt.Println(longest("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz") == "abcdefghijklmnopqrstuvwxyz")
}
