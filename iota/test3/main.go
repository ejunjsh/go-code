package main

import (
	"fmt"
)

const(
    a= 10+iota
    b
    c
)

const(
	d=1+iota
	e
	f
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
