package main

import (
	"fmt"
	"strings"
)

func removeSpaces(str string) string {
	characters := strings.Split(str, " ")
	return strings.Join(characters, "")
}

func main() {
	fmt.Println(removeSpaces("nubster is a kitty cat"))
}
