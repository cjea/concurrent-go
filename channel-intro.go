package main

import(
  "fmt"
  "time"
)

func main() {
  go printStuff()
  time.Sleep(time.Millisecond * 500)
  fmt.Println("All done!")
}

func printStuff() {
  for {
    fmt.Println("waiting . . .")
    time.Sleep(time.Millisecond * 50)
  }
}
