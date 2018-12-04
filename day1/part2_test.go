package day1

import "testing"

func TestProcess(t *testing.T) {
	sets := []struct{
		s []int
		r int
	} {
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	for _, d := range sets {
		if r := Process(d.s); r != d.r {
			t.Errorf("Failed %d gave %d expected %d", d.s, r, d.r)
		}
	}
}