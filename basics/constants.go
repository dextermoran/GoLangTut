package main

import "fmt"

const p string = "Death and Taxes"

const (
	Pi  = 3.14
	Pie = "cherry"

	P = iota // 1
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(PI)
	fmt.Println(Pie)
}
