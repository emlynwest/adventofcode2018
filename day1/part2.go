package day1

import "fmt"

// Part2 is the main entry point for day 1 part 2's solution.
func Part2() {
	fmt.Printf("Found: %d", Process(input))
}

// Process checks for duplicate values in the frequency list.
func Process(i []int) int {
	found := map[int]bool{
		// 0 is our starting value so it's always in there.
		0: true,
	}

	value := 0
	index := 0

	for ; ; {

		if index >= len(i) {
			index = 0
		}

		value = Alter(value, i[index])

		if found[value] {
			break
		}

		found[value] = true
		index++
	}

	return value
}