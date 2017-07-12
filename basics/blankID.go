package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
)

/*  (must import "log" as well)
func main() {
  res, err := http.Get("https://www.google.com")
  if err != nil {
    log.Fatal(err)
  }
  page, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", page)

}
*/

//the same thing w/ blank identifiers in place of err variables (no error handling)
func main() {
  res, _ := http.Get("https://www.google.com")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", page)
}
