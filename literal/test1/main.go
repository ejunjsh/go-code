package main

import "fmt"

func main() {
	a := 1
	b := 3

	fmt.Printf("%T\n", a / b)
	fmt.Printf("%T\n", a / 3)
	fmt.Printf("%T\n", a / 3.0)
	fmt.Printf("%T\n", 1 / 3)
	fmt.Printf("%T\n", 1 / 3.0)
}