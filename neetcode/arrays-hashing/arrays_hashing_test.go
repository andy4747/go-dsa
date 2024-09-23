package arrayshashing

import (
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
