package main

import "fmt"

type human interface{
	walk()
	run()
	eat()
}


type man struct{}

func (m *man) walk(){
	fmt.Println("man walk")
}

func (m *man) run(){
	fmt.Println("man run")
}

func (m *man) eat(){
	fmt.Println("man eat")
}

type woman struct{}

func (m *woman) walk(){
	fmt.Println("woman walk")
}

func (m *woman) run(){
	fmt.Println("woman run")
}

func (m *woman) eat(){
	fmt.Println("woman eat")
}

func main(){
	var h human=&man{}
	h.eat()
	h.run()
	h.walk()

	h=&woman{}
	h.eat()
	h.run()
	h.walk()
}