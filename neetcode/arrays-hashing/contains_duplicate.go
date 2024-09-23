package arrayshashing

func ContainsDuplicate(nums []int) bool {
	result := make(map[int]bool)
	for _, v := range nums {
		if _, exists := result[v]; exists {
			return true
		}
		result[v] = true
	}
	return false
}
