/*
Author: Angel Dhakal
File: inverted_right_pyramid.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-5-inverted-right-pyramid/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternInvertedRightPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
