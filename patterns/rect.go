/*
Author: Angel Dhakal
File: rect.go
Date: 2024-08-17
Problem: https://takeuforward.org/pattern/pattern-1-rectangular-star-pattern/
*/

package patterns

import (
	"fmt"
)

func PatternRectangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
