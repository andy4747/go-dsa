package arrays

import (
	"slices"
	"testing"
)

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		list     []int
		reversed []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{4, 5, 6}, []int{6, 5, 4}},
		{[]int{7, 8, 9, 10}, []int{10, 9, 8, 7}},
		{[]int{11, 12, 13, 14, 15}, []int{15, 14, 13, 12, 11}},
		{[]int{16, 17}, []int{17, 16}},
		{[]int{18}, []int{18}},
		{[]int{-1, -2, -3, -4}, []int{-4, -3, -2, -1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{100, 200, 300, 400}, []int{400, 300, 200, 100}},
		{[]int{123, 456, 789}, []int{789, 456, 123}},
		{[]int{10, 20, 30, 40, 50, 60}, []int{60, 50, 40, 30, 20, 10}},
		{[]int{5, 10, 15, 20, 25, 30, 35}, []int{35, 30, 25, 20, 15, 10, 5}},
	}
	for i, test := range tests {
		ReverseInplace(test.list)
		if !slices.Equal(test.list, test.reversed) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.reversed, test.list)
		}
	}
}
