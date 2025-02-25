package arrays

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		list   []int
		sorted []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{8, 7, 6, 5, 4}, []int{4, 5, 6, 7, 8}},
		{[]int{15, 16, 12, 11, 17, 10, 9}, []int{9, 10, 11, 12, 15, 16, 17}},
	}
	for i, test := range tests {
		BubbleSort(test.list)
		if slices.Equal(test.list, test.sorted) {
			t.Logf("Passed Test Case #%d. Got %#v", i, test.sorted)
		} else {
			t.Errorf("Failed Test Case #%d. Want %#v, got %#v", i, test.sorted, test.list)
		}
	}
}
