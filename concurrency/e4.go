package main

import (
  "log"
  "sync"
  "time"
)

type Message struct {
  Type   string
  Seq    int
  SendAt time.Time
}

func ping(ch chan Message, done chan bool, wg *sync.WaitGroup) {
  defer wg.Done()
  defer close(ch)
  exchanges := 3
  for i := 1; i <= exchanges; i++ {
    select {
    case <-done:
      return
    default:
      msg := Message{
        Type:   "ping",
        Seq:    i,
        SendAt: time.Now(),
      }
      log.Printf("Ping goroutine: Sending %s %d", msg.Type, msg.Seq)

      ch <- msg

      //wait and receivce for pong
      response, ok := <-ch
      if !ok {
        log.Printf("Ping goroutine: Channel closed")
        return
      }
      log.Printf("Ping goroutine: Received %s %d (round-trip: %v)",
        response.Type,
        response.Seq,
        time.Since(response.SendAt))

      // Add small delay between rounds
    }

  }
}

func pong(ch chan Message, wg *sync.WaitGroup) {
  defer wg.Done()

  for {
    // wait for ping
    msg, ok := <-ch
    if !ok {
      log.Printf("Pong goroutine: Channel closed")
      return
    }
    log.Printf("Pong goroutine: Received %s %d", msg.Type, msg.Seq)
    // Send pong
    response := Message{
      Type:   "pong",
      Seq:    msg.Seq,
      SendAt: time.Now(),
    }
    log.Printf("Pong goroutine: Sending %s %d", response.Type, response.Seq)
    ch <- response
  }
}

func main() {
  start := time.Now()
  ch := make(chan Message)
  done := make(chan bool)
  var wg sync.WaitGroup

  wg.Add(2)

  go ping(ch, done, &wg)
  go pong(ch, &wg)

  // Wait for completion
  wg.Wait()
  close(done)
  log.Printf("Conversation completed in %v", time.Since(start))
}