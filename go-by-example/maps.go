package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println("after deletion:", m)

	_, appears := m["k2"]
	fmt.Println(appears)

	_, appears = m["k1"]
	fmt.Println(appears)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
