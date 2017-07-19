package main

import "fmt"

func main() {
  data := []float64{42, 65, 76, 93, 35, 17}
  n := average(data...)
  fmt.Println(n)
}

func average(x ...float64) float64 {
  var total float64
  for _, d := range x {
    total += d
  }
  return total / float64(len(x))
}
