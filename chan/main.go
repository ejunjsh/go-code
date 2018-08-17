package main

import (
	"fmt"
	"time"
)

func main()  {
	c:=make(chan struct{},1)

	go func() {
		for {
			select {
				case <-c:
					fmt.Println("haha done")
					return
			}
		}
	}()

	go func() {
		for {
			select {
				case <-c:
					fmt.Println("xixi done")
					return
			}
		}
	}()

	time.Sleep(2*time.Second)

	c<- struct{}{}

	//close(c)

	time.Sleep(2*time.Second)
}
