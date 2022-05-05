package main

import "fmt"

func createRow(firstNum, rowNum int) []int {
	currentIdx := 0
	currentNum := firstNum
	row := make([]int, rowNum)

	for currentIdx < rowNum {
		row[currentIdx] = currentNum
		currentIdx += 1
		currentNum += 2
	}

	return row
}

func RowSumOddNumbers(n int) int {
	currentNum := 1
	currentRowNum := 1
	var nthRow []int

	for currentRowNum <= n {
		nthRow = createRow(currentNum, currentRowNum)
		currentNum = nthRow[len(nthRow)-1] + 2
		currentRowNum += 1
	}

	sum := 0
	for _, num := range nthRow {
		sum += num
	}

	return sum
}

func main() {
	// fmt.Println(createRow(1, 1)) // [1]
	// fmt.Println(createRow(3, 2)) // [3 5]
	// fmt.Println(createRow(7, 3)) // [7 9 11]

	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(3))
}
