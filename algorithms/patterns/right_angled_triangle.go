/*
Author: Angel Dhakal
File: right_angled_triangle.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-2-right-angled-triangle-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/

package patterns

import "fmt"

func PatternRightAngledTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
