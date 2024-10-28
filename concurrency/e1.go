package concurrency

import (
  "fmt"
  "sync"
)

/*
### E1. Simple Hello World Goroutine
- **Task:** Create a goroutine that prints "Hello, Concurrent World!"
- **Requirements:**
  - Main function should wait for goroutine to complete
  - Use appropriate synchronization mechanism
- **Learning Focus:** Basic goroutine creation and synchronization
- **Tips:** Consider using WaitGroup or channel for synchronization
- **Time:** 5 mins
*/

func greet(ch chan<- string) {
  ch <- "Hello, Concurrent World!"
}

func E1(){
  var wg sync.WaitGroup
  wg.Add(1)
  go func () {
    fmt.Println("Hello, Concurrent World!")
    wg.Done()
  }()
  wg.Wait()
}

func main(){
  ch := make(chan string)
  go greet(ch)
  fmt.Println(<-ch)
  E1()
}
