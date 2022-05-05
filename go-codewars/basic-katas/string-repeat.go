package main

func repeatStr(repeater int, str string) string {
	repeated := ""

	for count := 0; count < repeater; count++ {
		repeated += str
	}

	return repeated
}
