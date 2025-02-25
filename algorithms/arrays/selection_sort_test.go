package arrays

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		list     []int
		reversed []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{6, 5, 4}, []int{4, 5, 6}},
	}
	for i, test := range tests {
		SelectionSort(test.list)
		if slices.Equal(test.list, test.reversed) {
			t.Logf("Passed Test Case #%d. Got %#v", i, test.reversed)
		} else {
			t.Errorf("Failed Test case #%d. Want %#v, got %#v", i, test.reversed, test.list)
		}
	}
}
