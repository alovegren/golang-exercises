package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
	// we dereference the pointer from its memory address to the *value* at that address
	// we then have a dereferenced pointer
	// we assign a new value to that dereferenced pointer
	// this causes the value at the same memory address as before to now be occupied by the new value
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	// the 'address of' operator passes the memory address of i to the function
	// in other words, it allows us to pass a pointer to zeroptr
	fmt.Println("zeroval:", i)

	fmt.Println("pointer:", &i)
}
