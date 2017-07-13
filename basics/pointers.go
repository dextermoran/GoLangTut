package main

import "fmt"

/*

func main() {
   a := 23

   fmt.Println("a - ", a)
   fmt.Println("a's mem address", &a)
}
*/

const metersToYards float64 = 1.09361

func main() {
  var meters float64
  fmt.Println("meters swam: ")
  fmt.Scan(&meters)
  yards := meters * metersToYards
  fmt.Println(meters, "meters is", yards, " Yards")
}
