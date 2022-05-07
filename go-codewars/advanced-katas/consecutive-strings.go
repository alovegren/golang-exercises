package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) (longestString string) {
	var longestLength int

	for startIdx := 0; startIdx+k <= len(strarr); startIdx++ {
		stringCandidate := strings.Join(strarr[startIdx:startIdx+k], "")
		candidateLength := len(stringCandidate)

		if candidateLength > longestLength {
			longestLength = candidateLength
			longestString = stringCandidate
		}
	}

	return
}

/*
We want the longest sequence that can be formed by k *consecutive* strings.

So we basically want to form a sliding window of n many strings, stopping when the length of that window would cause us to reach the end of the slice. We can greedily dub subsequences as longer and longer and return the winner.
*/

func main() {
	fmt.Println(LongestConsec([]string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}, 2))
}
