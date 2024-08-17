/*
Author: Angel Dhakal
File: right_angled_number_pyramid.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-3-right-angled-number-pyramid/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternRightAngledNumberPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
