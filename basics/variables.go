package main

import "fmt"

var MyName = "Dexter"

func main() {
  var lastName string
  lastName = "moran"

  var fullname string = "Dexter moran"

  shrthnd := 111
  shrthnd2 := "oneoneone"

  fmt.Printf("%T \n", shrthnd)
  fmt.Printf("%T \n", shrthnd2)

  fmt.Println(MyName + lastName)
  fmt.Println(fullname)
}
