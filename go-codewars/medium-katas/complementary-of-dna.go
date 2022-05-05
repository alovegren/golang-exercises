package main

import (
	"fmt"
	"strings"
)

func dnaStrand(dna string) (complements string) {
	nucleobases := strings.Split(dna, "")

	for _, nucleobase := range nucleobases {
		switch nucleobase {
		case "A":
			complements += "T"
		case "T":
			complements += "A"
		case "G":
			complements += "C"
		case "C":
			complements += "G"
		}
	}

	return complements
}

func main() {
	fmt.Println(dnaStrand(""))     // ""
	fmt.Println(dnaStrand("ATGC")) // TACG
	fmt.Println(dnaStrand("GTAT")) // CATA
	fmt.Println(dnaStrand("AAAA")) // TTTT
}
