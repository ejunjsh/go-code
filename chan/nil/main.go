package main

import "fmt"

func main() {

	var nilChan chan struct{}

	select {
	//nil chan always block
	case <-nilChan:
	default:
		fmt.Println("always here")
	}
}
