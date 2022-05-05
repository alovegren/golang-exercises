package main

import (
	"fmt"
)

func PrinterError(s string) string {
	errorCount := 0
	runeValueM := 109
	// iterate thru letters in string
	for _, runeValue := range s {
		if int(runeValue) > runeValueM {
			errorCount += 1
		}
	}

	return fmt.Sprintf("%s/%s", fmt.Sprint(errorCount), fmt.Sprint(len(s)))
}

func main() {
	fmt.Println(PrinterError("aaabbbbhaijjjm"))         // 0/14
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm")) // 8/22
}
