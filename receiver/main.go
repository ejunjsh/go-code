package main

import "fmt"

type person struct {
	age int
}


func (p *person) add(){
	p.age++
}

func (p person) reduce(){
	p.age--
}

func main()  {
	p:=person{age:0}
	p.add()
	p.reduce()
	fmt.Printf("person age is %d\n",p.age)

	p2:=&person{age:0}
	p2.add()
	p2.reduce()
	fmt.Printf("person age is %d\n",p2.age)
}
