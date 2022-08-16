package main

import "fmt"

type Person struct{
	firstname string
	lastname string
}

func changeMe(p *Person){
	(*p).firstname="Ola"
	(*p).lastname="Alo"
}

func (p* Person) setLastName(ln string){
	(*p).lastname=ln
}

func main() {
	p1 := Person{}
	changeMe(&p1)
	fmt.Println(p1.firstname, p1.lastname)
	p1.setLastName("Pastel")
	fmt.Println(p1.firstname, p1.lastname)
}

