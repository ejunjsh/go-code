package main

import "fmt"

func merge(out chan<- int, a, b <-chan int) {
	for {
		select {
		case v := <-a:
			out <- v
		case v := <- b:
			out <- v
		}
	}
}

// always to output zero
func main(){
	out,a,b:=make(chan int),make(chan int),make(chan int)

	go func() {
		merge(out,a,b)
	}()

	close(a)
	close(b)

	for {
		select{
		case v := <-out:
			fmt.Println(v)
		}
	}

}