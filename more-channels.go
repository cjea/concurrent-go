package main

import (
  "fmt"
  "time"
) 

func main() {
  c := make(chan string)
  go ping(c)
  go pong(c)
  go reader(c)

  time.Sleep(time.Millisecond * 50)
  fmt.Println(<-c)
}

func ping(c chan string) {
  for i:=1; ;i++ {
    c <- "Ping!" 
  }
}

func pong (c chan string) {
  for i:=1; ;i++ {
    c <- "Pong!"
  }
}

func reader (c chan string) {
  for {
    message := <-c
    fmt.Println(message)
    time.Sleep(time.Millisecond*5)
  }
}
