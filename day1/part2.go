package day1

import "fmt"

func Part2() {
	fmt.Printf("Found: %d", Process(input))
}

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