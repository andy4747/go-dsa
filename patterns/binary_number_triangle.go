/*
Author: Angel Dhakal
File: binary_number_triangle.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-11-binary-number-triangle-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternBinaryNumberTriangle(n int) {
	start := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", start)
			start = 1 - start
		}
		fmt.Println()
	}
}
