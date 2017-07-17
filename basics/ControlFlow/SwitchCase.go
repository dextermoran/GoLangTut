package main

import "fmt"

func main() {
  switch "thing2" {
  case "thing0":
    fmt.Println("thing0")
  case "thing1":
    fmt.Println("thing1")
  default:
    fmt.Println("nothing")
  }
}
