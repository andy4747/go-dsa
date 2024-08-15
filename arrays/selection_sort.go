/*
Author: Angel Dhakal
Date: 2024-08-14
File: selection_sort.go

Problem: Selection sort the given array
Pattern {
Iterate through the array, compare the ith element from unsorted portion,with the remaining elements and swap the min value to that ith element.
Sorted portion will be swapped at the start of the array

Time Complexity: O(n^2)
Space Complexity: O(1)
}
*/
package arrays

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
}
