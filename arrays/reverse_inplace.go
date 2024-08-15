/*
Author: Angel Dhakal
File: reverse_inplace.go
Date: 2024-08-08

Problem: Reverse the given array in place.
Pattern{
make two pointers that iterates the array from left and right till left is smaller than right, and on every iteration
swap the element at the left and rigt pointers.

Time Complexity: O(n)
Space Complexity: O(1)
}
*/
package arrays

func ReverseInplace(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
