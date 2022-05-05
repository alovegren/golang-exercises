package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CountBits(integer uint) (oneCount int) {
	binaryInt := strconv.FormatInt(int64(integer), 2)
	bits := strings.Split(binaryInt, "")

	for _, bit := range bits {
		if bit == "1" {
			oneCount++
		}
	}

	return
}

func main() {
	fmt.Println(CountBits(1234))
}
