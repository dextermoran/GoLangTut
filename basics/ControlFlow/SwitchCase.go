package main

import "fmt"

func main() {
	switch "thing0" {
	case "thing0":
		fmt.Println("thing0")
		fallthrough
	case "thing1":
		fmt.Println("thing1")
	default:
		fmt.Println("nothing")
	}
}
