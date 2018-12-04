package day2

import "testing"

func TestGenerateFrequencies(t *testing.T) {

	sets := []struct{
		s string
		r Frequencies
	} {
		{"abcdef", Frequencies{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1}},
		{"bababc", Frequencies{'b': 3, 'a': 2, 'c': 1}},
		{"abbcde", Frequencies{'a': 1, 'b': 2, 'c': 1, 'd': 1, 'e': 1}},
		{"abcccd", Frequencies{'a': 1, 'b': 1, 'c': 3, 'd': 1}},
		{"aabcdd", Frequencies{'a': 2, 'b': 1, 'c': 1, 'd': 2}},
		{"abcdee", Frequencies{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 2}},
		{"ababab", Frequencies{'a': 3, 'b': 3}},
	}

	for _, d := range sets {
		f := GenerateFrequencies(d.s)

		for k, v := range d.r {
			if v != f[k] {
				t.Errorf("%s produced incorrect frequency for %s. Got %d want %d", d.s, string(k), v, d.r[k])
			}
		}
	}
}

func TestCounts(t *testing.T) {

	sets := []struct{
		s []string
		twos int
		threes int
	} {
		{[]string{"abcdef"}, 0, 0},
		{[]string{"bababc"}, 1, 1},
		{[]string{"abbcde"}, 1, 0},
		{[]string{"abcccd"}, 0, 1},
		{[]string{"aabcdd"}, 1, 0},
		{[]string{"abcdee"}, 1, 0},
		{[]string{"ababab"}, 0, 1},
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 4, 3},
	}

	for _, d := range sets {
		twos, threes := Counts(d.s)

		if twos != d.twos {
			t.Errorf("Incorrect twos count for %s. Want %d, got %d", d.s, d.twos, twos)
		}

		if threes != d.threes {
			t.Errorf("Incorrect threes count for %s. Want %d, got %d", d.s, d.threes, threes)
		}
	}
}
