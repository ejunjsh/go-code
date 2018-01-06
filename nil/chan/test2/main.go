package main

import "fmt"

func merge(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				fmt.Println("a is nil")
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				fmt.Println("b is nil")
				continue
			}
			out <- v
		}
	}
	fmt.Println("close out")
	close(out)
}

func main(){
	out,a,b:=make(chan int),make(chan int),make(chan int)

	go func() {
		merge(out,a,b)
	}()

	close(a)
	close(b)

	for out!=nil{
		select{
		case v,ok:= <-out:
			if !ok{
				out=nil
				fmt.Println("out is nil")
			}
			fmt.Println(v)
		}
	}

}