package main

import "fmt"

func main() {

  //anonymous function
  greeting := func() {
    fmt.Println("Hello, World")
  }

  greeting()
}
