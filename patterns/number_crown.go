/*
Author: Angel Dhakal
File: number_crown.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-12-number-crown-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternNumberCrown(n int) {
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d", j)
		}
		for j := 1; j < space+1; j++ {
			fmt.Printf(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Printf("%d", j)
		}
		fmt.Println()
		space -= 2
	}
}
