/*
Author: Angel Dhakal
Date: 2024-08-14
File: contains_duplicate.go
Time Complexity: O(n)
Space Complexity: O(n)
Problem: LeetCode 217 - Contains Duplicate
Solution: {
Use a hash map to keep track of seen numbers. Iterate through the array once,
checking if each number is already in the map. If it is, return true (duplicate found).
If not, add the number to the map. If we finish iterating without finding a duplicate, return false.
}
*/
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
