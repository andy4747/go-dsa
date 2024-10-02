package arrayshashing

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, v := range nums {
		for j, w := range nums {
			if i != j {
				sum := v + w
				if sum == target {
					result = append(result, i, j)
					return result
				}
			}
		}
	}
	return result
}
