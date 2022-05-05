package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals() // this is called multiple assignment
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
