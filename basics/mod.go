package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		x := i % 3
		y := i % 5
		if x == 0 && y == 0 {
			fmt.Println("FizzBuzz")
			fmt.Println(i)
		} else if x == 0 {
			fmt.Println("fizz")
			fmt.Println(i)
		} else if y == 0 {
			fmt.Println("Buzz")
			fmt.Println(i)
		}
	}
}
