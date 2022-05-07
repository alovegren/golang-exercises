package main

import (
	"strings"
)

func DecodeMorse(morseCode string) string {
	morseWords := strings.Split(morseCode, "   ")
	decodeWords := make([]string, len(morseWords))

	for wordIdx, morseWord := range morseWords {
		decodeWord := make([]string, len(morseWord))
		morseChars := strings.Split(morseWord, " ")

		for charIdx, morseChar := range morseChars {
			decodeWord[charIdx] = MORSE_CODE[morseChar]
		}

		decodeWords[wordIdx] = strings.Join(decodeWord, "")
	}

	return strings.Join(decodeWords, " ")
}

func main() {
	DecodeMorse(".... . -.--   .--- ..- -.. .")
}
