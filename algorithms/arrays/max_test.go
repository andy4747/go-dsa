package arrays

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		list []int
		max  int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{5, 6, 7}, 7},
		{[]int{9, 10, 15}, 15},
	}
	for i, test := range tests {
		max := Max(test.list)
		if max == test.max {
			t.Logf("Passed Case #%d. Got %d", i, test.max)
		} else {
			t.Fatalf("Failed Case #%d. Want %d, got %d", i, test.max, max)
		}
	}
}
