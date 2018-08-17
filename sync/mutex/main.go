package main

import (
	"sync"
	"fmt"
	"time"
)

type lock struct {
	l sync.Mutex
}

func main(){
	l:=lock{}

	go func() {
		l.l.Lock()
		defer l.l.Unlock()
		fmt.Println("haha get the lock")
		time.Sleep(1000*time.Second)
	}()

	go func() {
		l.l.Lock()
		defer l.l.Unlock()
		fmt.Println("xixi get the lock")
		time.Sleep(1000*time.Second)
	}()

	time.Sleep(1000*time.Second)
}
