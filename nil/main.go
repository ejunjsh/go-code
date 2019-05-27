package main

import "fmt"

func main() {
	var a int
	var b bool
	var g string
	var c func()
	var d []int
	var e map[string]int
	var f [10]string
	var h interface{}
	var i error
	var j chan int
	var k *int

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
	fmt.Println(d == nil)
	fmt.Println(e == nil)
	fmt.Println(h == nil)
	fmt.Println(i == nil)
	fmt.Println(j == nil)
	fmt.Println(k == nil)
}

//output:
//0 false <nil> [] map[] [         ]  <nil> <nil> <nil> <nil>
//true
//true
//true
//true
//true
//true
