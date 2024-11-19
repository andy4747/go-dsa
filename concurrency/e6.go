package concurrency

import (
	"fmt"
	"sync"
)

/*
### E6. Simple Producer (1-10)
- **Task:** Implement a producer that generates numbers 1-10
- **Requirements:**
  - Use a goroutine as producer
  - Send numbers to a channel
  - Main function receives and prints
  - Close channel when done
- **Learning Focus:** Channel closing, range over channel
- **Time:** 5 mins
*/

func produce_numbers(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func MainE6() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go produce_numbers(ch, &wg)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}
