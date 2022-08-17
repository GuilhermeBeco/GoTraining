package main

import (
	"fmt"
)

type Person struct{
	first string
	last string
	age int
}

type Human interface{
	speak()
}

func (p *Person) speak(){
	fmt.Println(p.first, p.last, p.age)
}

func saySomething(h Human){
	h.speak()
}

func main(){
	p1 := Person{
		first:"Guilherme",
		last:"Beco",
		age:23,
	}
	//saySomething(p1) --- Does not work
	saySomething(&p1)
}