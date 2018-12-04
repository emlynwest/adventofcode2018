package day1

import "fmt"

// Part1 entry point for day 1, task 1
func Part1() {
	r := Sequence(0, input)

	fmt.Printf("Result: %d", r)
}

// Alter performs a modification on the input value
func Alter(input int, change int) int {
	return input + change
}

// Sequence applies numeric transformations to start value
func Sequence(start int, values []int) int {

	a := start

	for _, i := range values {
		a = Alter(a, i)
	}

	return a
}
