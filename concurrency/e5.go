package concurrency

import (
	"fmt"
)

/*
### E5. Number Squaring with Channels
- **Task:** Create a number squaring service using goroutine
- **Requirements:**
  - Input channel receives numbers
  - Goroutine squares each number
  - Output channel returns results
  - Process 5 numbers
- **Learning Focus:** Channel direction, data processing
*/

func square(inp_ch chan int, out_ch chan int) {
	for num := range inp_ch {
		out_ch <- num * num
	}
	close(out_ch)
}

func Main_e5() {
	inp_ch := make(chan int)
	out_ch := make(chan int)

	nums := []int{2, 3, 4, 5, 6}

	go square(inp_ch, out_ch)

	go func() {
		for _, v := range nums {
			inp_ch <- v
		}
		close(inp_ch)
	}()

	for v := range out_ch {
		fmt.Println(v)
	}

	fmt.Println("Program Completed")
}
