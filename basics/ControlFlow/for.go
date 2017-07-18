package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for x := 0; x < 10; x++ {
			fmt.Println(i, ".", x)
		}
	}
}
