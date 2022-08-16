package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age int
}

func (p Person) speak(){
	fmt.Println("My name is",p.firstName,p.lastName,"and i'm",p.age,"years old")
}

func main() {
	p1 := Person{
		firstName: "Guilherme",
		lastName:  "Beco",
		age: 23,
	}

	p1.speak()
}
