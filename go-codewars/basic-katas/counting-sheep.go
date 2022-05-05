package main

func countingSheep(sheep []bool) (total int) {
	for idx := range sheep {
		if sheep[idx] {
			total += 1
		}
	}

	return total
}
