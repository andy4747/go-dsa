package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		list          []int
		searchElement int
		index         int
	}{
		// Basic cases
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},
		{[]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 11, 7},
		{[]int{0, 10, 20, 35, 45, 55, 64, 74, 84, 99}, 99, 9},

		// Edge cases
		{[]int{1, 2, 3, 4, 5}, 1, 0},  // Element at the start
		{[]int{1, 2, 3, 4, 5}, 5, 4},  // Element at the end
		{[]int{1, 2, 3, 4, 5}, 3, 2},  // Element in the middle
		{[]int{1, 2, 3, 4, 5}, 0, -1}, // Element not in the list (less than min)
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Element not in the list (greater than max)
		{[]int{}, 1, -1},              // Empty list
		{[]int{5}, 5, 0},              // Single element list, element found
		{[]int{5}, 4, -1},             // Single element list, element not found

		// Larger lists
		{[]int{-100, -50, -10, 0, 10, 20, 30, 40, 50, 60}, 30, 6},  // Larger list, positive numbers
		{[]int{-100, -50, -10, 0, 10, 20, 30, 40, 50, 60}, -50, 1}, // Larger list, negative numbers

		// List with negative and positive numbers
		{[]int{-10, -5, 0, 5, 10}, -5, 1}, // Mixed sign list, negative number
		{[]int{-10, -5, 0, 5, 10}, 5, 3},  // Mixed sign list, positive number

		{[]int{1, 2, 3, 4, 5}, 7, -1},      // Element not in list, larger than max
		{[]int{-5, -4, -3, -2, -1}, -3, 2}, // List of all negatives
		{[]int{-5, -4, -3, -2, -1}, 0, -1}, // Element not in list, positive number in negative list
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
