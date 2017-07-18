package main

import "fmt"

func main() {
	fmt.Println(funko("eleven"))

	n := average(22, 33, 44, 100)
	fmt.Println(n)
}

func funko(name string) string {
	return fmt.Sprint("upside down ", name)
}

func average(x ...float64) float64 {
	fmt.Println(x)
	fmt.Printf("%T", x)
	var total float64

	for _, v := range x {
		total += v
	}

	return total / float64(len(x))
}
