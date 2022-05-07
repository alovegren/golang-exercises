package main

/*
Rules

We can keep track of whether a brace is balanced by counting.

- We can advance counts up each time an opening brace is encountered, and down each time a closing brace is encountered
- By the time we've traversed through the whole string, the counts should be zero. Such an approach would work just fine for one type of brace, but we need extra rules now that there are three.

e.g. [(]) --> this guy is false.
based on our idea we'd have a map something like
idx: 0 --> { squareBrackets: 1, parentheses: 0 }
idx: 1 --> { squareBrackets: 1, parentheses: 1 }
idx: 2 --> { squareBrackets: 0, parentheses: 1 } <-- we know at this point we've failed

mb data structure should be parent/child-based?
or ordered in an array?

e.g. [(])
what if we push the glyph on when we encounter an open and pop it off when we encounter a close?
idx: 0 ["["]
idx: 1 ["[", "("]
idx: 2 ["[", "("] <-- there's no square bracket to pop off, so this is false

e.g.(){}[]
idx: 0 ["("]
idx: 1 [] since we have a ")" and our last element was a "(", we can pop it off and keep going
idx: 2 ["{"]
idx: 3 ["["]
idx: 4 []
we return true since the array is empty

([{}])
0 [ "(" ]
1 [ "(", "[" ]
2 [ "(", "[", "{"]
3 [ "(", "[" ] (legal pop)
4 [ "(" ] (legal pop)
5 [] (legal pop)
empty --> true

Algorithm
- Initialize a slice that can grow/shrink
- Create a map, mapping closing braces to their open counterparts
- Iterate through the string
	- If the current character is an opening brace, append it to the slice
	- Otherwise, if the value of the key in the map equivalent to the current character is equal to the last element in the slice,
		- Shrink the slice by one ("popping off" the last element)
	- Otherwise,
		- Return false
- If, at the end of the iteration, the array is empty, return true
- Otherwise, return false
*/

import (
	"fmt"
	"strings"
)

/*
e.g. ]}


*/

func ValidBraces(str string) bool {
	originalBraces := strings.Split(str, "")
	braces := []string{}
	closedToOpenBraces := map[string]string{")": "(", "]": "[", "}": "{"}

	for _, brace := range originalBraces {
		switch {
		case brace == "(", brace == "[", brace == "{":
			braces = append(braces, brace)
		case len(braces) >= 1: // it's a closing brace and there is at least one opening brace in the braces arr
			lastOpeningBrace := braces[len(braces)-1]

			if closedToOpenBraces[brace] == lastOpeningBrace {
				braces = braces[:len(braces)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	if len(braces) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(ValidBraces("(){}[]") == true)
	fmt.Println(ValidBraces("([{}])") == true)
	fmt.Println(ValidBraces("(}") == false)
	fmt.Println(ValidBraces("[(])") == false)
	fmt.Println(ValidBraces("[({})](]") == false)
}
