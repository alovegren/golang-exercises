package main

import "fmt"

func MaxMultiple(d, b int) int {
	for n := b; n > 0; n-- {
		if n%d == 0 {
			return n
		}
	}

	return 1
}

func main() {
	fmt.Println(MaxMultiple(2, 7))
}
