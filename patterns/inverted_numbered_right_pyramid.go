/*
Author: Angel Dhakal
File: inverted_numbered_right_pyramid.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-6-inverted-numbered-right-pyramid/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternInvertedNumberedRightPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
