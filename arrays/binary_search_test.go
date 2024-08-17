package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		list          []int
		searchElement int
		index         int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},
		{[]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 11, 7},
		{[]int{0, 10, 20, 35, 45, 55, 64, 74, 84, 99}, 99, 9},
	}
	for i, test := range tests {
		searchedIndex := BinarySearch(test.list, test.searchElement)
		if searchedIndex == test.index {
			t.Logf("Passed Test Case #%d. Got Index: %d", i, test.index)
		} else {
			t.Errorf("Failed Test Case #%d. Want Index: %d, Got Index: %d", i, test.index, searchedIndex)
		}
	}
}
