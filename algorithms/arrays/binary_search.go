/*
Author: Angel Dhakal
File: binary_search.go
Date: 2024.08-16
Time Complexity: O(log n)
Space Complexity: O(1)
Problem: binary search on an array
Solution:{
	You should be able to do a binary search without a note.
}
*/

package arrays

func BinarySearch(arr []int, searchElement int) int {
	first := 0
	last := len(arr) - 1
	for first <= last {
		mid := first + (last-first)/2
		if arr[mid] == searchElement {
			return mid
		}
		if searchElement > arr[mid] {
			first = mid + 1
		}
		if searchElement < arr[mid] {
			last = mid - 1
		}
	}
	return -1
}
