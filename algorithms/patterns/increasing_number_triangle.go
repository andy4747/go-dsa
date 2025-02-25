/*
Author: Angel Dhakal
File: increasing_number_triangle.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-13-increasing-number-triangle-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternIncreasingNumberTriangle(n int) {
	lastNum := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", lastNum)
			lastNum = lastNum + 1
		}
		fmt.Println()
	}
}
