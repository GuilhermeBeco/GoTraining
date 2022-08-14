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

	map1 := map[string]person{}
	map1[p1.lastName] = p1
	for _, t := range map1 {
		for _, v := range t.iceCream {
			fmt.Println(t.firstName, t.lastName, v)
		}
	}
}
