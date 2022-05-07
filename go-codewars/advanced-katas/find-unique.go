package main

import (
	"fmt"
)

func determineUniq(slice []float32) (value float32, found bool) {
	switch {
	case len(slice) < 3:
		return
	case (slice[0] == slice[1] && slice[1] == slice[2]):
		return
	default:
		found = true
		switch {
		case slice[0] == slice[1]:
			value = slice[2]
		case slice[0] == slice[2]:
			value = slice[1]
		default:
			value = slice[0]
		}
	}

	return
}

func FindUniq(slice []float32) float32 {
	for idx := 0; idx < len(slice); idx = idx + 3 {
		var currentSlice []float32

		if idx > len(slice)-3 {
			currentSlice = slice[idx:]
		} else {
			currentSlice = slice[idx : idx+3]
		}

		if value, foundUnique := determineUniq(currentSlice); foundUnique {
			return value
		}
	}

	lastThree := slice[len(slice)-3:]
	uniqMember, _ := determineUniq(lastThree)

	return uniqMember
}

func main() {
	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}))
}

/*
We essentially want a performant uniq method, knowing that we only have one unique member.

We could progressively divide the slices... So long as we have three numbers in the slice, we can determine from that slice which the unique value is. So our base case is a group of three... but our numbers aren't necessarily in any order, so what good does splitting them do?

Some cases
The unique number is in one of the first three positions
The unique number is in one of the last three positions
The unique number is somewhere in the middle
*/
