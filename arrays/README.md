# Reverse an Array In-Place

Easy | Topics | Companies | Hint

Given an array of integers `nums`, reverse the array in-place. An in-place algorithm modifies the input directly, using a constant amount of extra memory.

You should modify the input array and not return anything.

Example 1:

Input: nums = [1,2,3,4,5]
Output: [5,4,3,2,1]
Explanation: The array is reversed in-place.

Example 2:

Input: nums = [8,9,10,11]
Output: [11,10,9,8]

Example 3:

Input: nums = [1]
Output: [1]

Constraints:

- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1

Follow up: Can you solve it with O(1) extra space complexity?


# Remove Duplicates from Sorted Array

## Problem Statement

Given a sorted array of integers, remove the duplicates in-place such that each element appears only once and return the new length.

Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.

## Example 1:

**Input:** nums = [1,1,2]
**Output:** 2, nums = [1,2,_]
**Explanation:** Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively. It does not matter what you leave beyond the returned k (hence they are underscores).

## Example 2:

**Input:** nums = [0,0,1,1,1,2,2,3,3,4]
**Output:** 5, nums = [0,1,2,3,4,_,_,_,_,_]
**Explanation:** Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively. It does not matter what you leave beyond the returned k (hence they are underscores).

## Constraints:

- 1 <= nums.length <= 3 * 10^4
- -100 <= nums[i] <= 100
- nums is sorted in non-decreasing order.

## Function Signature:

```go
func removeDuplicates(nums []int) int
```
