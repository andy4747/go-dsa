/*
Author: Angel Dhakal
File: max.go
Date: 2024-08-08
Time Complexity: O(n)
Space Complexity: O(1)
Problem: return the max int from the array of int.
Solution:{
}
*/
package arrays

// Max returns max number from the given slice
func Max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
