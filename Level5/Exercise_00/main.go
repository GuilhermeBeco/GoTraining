package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Guilherme",
		lastName:  "Beco",
		iceCream:  []string{"Mint", "Caramel"},
	}

	for _, v := range p1.iceCream {
		fmt.Println(p1.firstName, p1.lastName, v)
	}
}
