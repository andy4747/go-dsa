/*
Author: Angel Dhakal
File: right_angled_number_pyramid2.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-4-right-angled-number-pyramid-ii/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternRightAngledNumberPyramid2(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
}
