package main

import (
	"fmt"
	"strings"
)

var print = fmt.Println

func main() {
	print("Contains: ", strings.Contains("test", "es"))
	print("Count:    ", strings.Count("test", "t"))
	print("HasPrefix:", strings.HasPrefix("test", "te"))
	print("HasSuffix ", strings.HasSuffix("test", "st"))
	print("Index:    ", strings.Index("test", "e"))
	print("Join:     ", strings.Join([]string{"a", "b"}, "-"))
	print("Repeat:   ", strings.Repeat("a", 5))
	print("Replace:  ", strings.Replace("foo", "o", "0", -1))
	print("Replace:  ", strings.Replace("foo", "o", "0", 1))
	print("Split:    ", strings.Split("a-b-c-d-e", "-"))
	print("ToLower:  ", strings.ToLower("KITTENS"))
	print("ToUpper:  ", strings.ToUpper("kittens "))
}
