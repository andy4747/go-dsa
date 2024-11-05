package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
 E2. WaitGroup with 3 Goroutines
- **Task:** Launch 3 goroutines that each sleep for different durations and print their IDs
- **Requirements:**
  - Use sync.WaitGroup for synchronization
  - Each goroutine sleeps for (ID * 100)ms
  - Main function must wait for all goroutines to finish
- **Learning Focus:** WaitGroup usage pattern
- **Time:** 5 mins
*/

func sleep(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Println(id)
}

func Main_e2() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sleep(i, &wg)
	}
	wg.Wait()
}
