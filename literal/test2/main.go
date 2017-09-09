package main

import "fmt"

func main() {
	a := 1

	fmt.Println(a / 3)
	fmt.Println(a / 3.0)
	fmt.Println(a / 3.1) //类型错误
}