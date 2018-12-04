package day1

import "testing"

func TestAlter(t *testing.T) {
	sets := []struct{
		a int
		b int
		r int
	}{
		{1, 1, 2},
		{0, 1, 1},
		{0, -1, -1},
	}

	for _, d := range sets {
		if r := Alter(d.a, d.b); r != d.r {
			t.Errorf("Failed with %d %d = %d got %d", d.a, d.b, d.r, r)
		}
	}
}

func TestSequence(t *testing.T) {

	sets := []struct{
		s int
		t []int
		r int
	} {
		{0,[]int{+1, +1, +1}, 3},
		{0,[]int{+1, +1, -2}, 0},
		{0,[]int{-1, -2, -3}, -6},

	}

	for _, d := range sets {
		if r := Sequence(d.s, d.t); r != d.r {
			t.Errorf("Expected %d, got %d start: %d set: %d", d.r, r, d.s, d.t)
		}
	}
}