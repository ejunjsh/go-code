package main

import (
	"fmt"
)

const(
    a= 10+iota
    b
    c
    d
	e=1+iota
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
