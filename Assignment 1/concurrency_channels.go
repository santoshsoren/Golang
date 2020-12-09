package main

import (
  "fmt"
  "time"
)

func kloudone(c chan string) {
  for i := 0; i < 5; i++ {
    c <- "kloudone"
  }
}

func display(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)
  fmt.Println(display(c))

  go kloudone(c)
  go kloudone(c)

  var input string
  fmt.Scanln(&input)
}