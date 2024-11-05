package concurrency

/*
 E3. Buffered Channel (size 5)
- **Task:** Implement a number generator using buffered channel
- **Requirements:**
  - Create a buffered channel of size 5
  - Send numbers 1-10 through it
  - Demonstrate channel blocking when buffer is full
  - Print each received number
- **Learning Focus:** Buffered channel behavior
- **Time:** 5 mins
*/

func generate(limit int, ch chan<- int) {
	for i := 1; i <= limit; i++ {
		ch <- i
		println("sent: ", i)
	}
	close(ch)
}

func Main_e3() {
	ch := make(chan int, 5)
	go generate(10, ch)
	for n := range ch {
		println("rec: ", n)
	}
}
