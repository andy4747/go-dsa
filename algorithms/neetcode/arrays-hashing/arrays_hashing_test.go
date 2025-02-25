package arrayshashing

import (
	"slices"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	t.Run("ContainsDuplicate", func(t *testing.T) {
		nums := []int{10, 20, 20, 30}
		nums2 := []int{10, 20, 30}
		res1 := ContainsDuplicate(nums)
		res2 := ContainsDuplicate(nums2)
		if res1 == true {
			t.Logf("Test Passed. Want true Got %v", res1)
		} else {
			t.Errorf("Test Failed. Want true Got %v", res1)
		}
		if res2 == false {
			t.Logf("Test Passed. Want true Got %v", res2)
		} else {
			t.Errorf("Test Failed. Want true Got %v", res2)
		}
	})
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, test := range tests {
		got := twoSum(test.nums, test.target)
		if slices.Equal(got, test.result) {
			t.Logf("Passed Test Case #%d. Got %#v", i, test.result)
		} else {
			t.Errorf("Failed Test Case #%d. Want %#v, got %#v", i, test.result, got)
		}
	}
}
