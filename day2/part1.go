package day2

import (
	"fmt"
)

type Frequencies map[rune]int

// Part1 runs day 2 part 1's solution
func Part1() {
	twos, threes := Counts(input)

	fmt.Printf("Result: %d", twos * threes)
}

// GenerateFrequencies generates frequency for each character in a string
func GenerateFrequencies(s string) Frequencies {
	f := Frequencies{}

	for _, c := range s {
		f[c]++
	}

	return f
}

// Counts loops over them and add to the 2 and 3 counts if correct counts exist
func Counts(s []string) (twos, threes int) {

	for _, v := range s {
		var added2, added3 bool

		for _, c := range GenerateFrequencies(v) {
			if c == 2 && !added2{
				twos++
				added2 = true
			}

			if c == 3 && !added3 {
				threes++
				added3 = true
			}
		}
	}

	return twos, threes
}
