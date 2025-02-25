package arrays

import (
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		list   []int
		sorted []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{9, 8, 7, 6, 5}, []int{5, 6, 7, 8, 9}},
	}
	for i, test := range tests {
		InsertionSort(test.list)
		if slices.Equal(test.list, test.sorted) {
			t.Logf("Passed Test Case #%d. Got %#v.", i, test.sorted)
		} else {
			t.Errorf("Failed Test Case #%d. Want %#v, Got %#v.", i, test.sorted, test.list)
		}
	}
}
