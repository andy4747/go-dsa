package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		ch <- rand.Intn(100) + 1
	}
	close(ch)
	wg.Done()
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
	}
	wg.Done()
}

func MainE7() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
