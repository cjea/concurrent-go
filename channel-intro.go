package main

import "fmt"
import "strconv"

func main() {
  fmt.Println("Entering main . . . ")
  
  done := make(chan bool)
  go write_to_channel(done)

  /* trying to read from this channel will block 
    main's execution until our goroutine writes a value to the channel 
    */
  <-done

  fmt.Println("Back in main")
  fmt.Println("All done!")
}

func write_to_channel(some_channel chan bool) {
  for i := 1; i < 6; i++ {
    fmt.Println("Goroutine rep " + strconv.Itoa(i))
  }
  // time to write to the channel that was passed in to the function
  some_channel <- true
}
