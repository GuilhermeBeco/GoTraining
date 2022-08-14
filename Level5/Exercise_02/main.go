package main

import "fmt"

type Car struct {
	doors int
	color string
}

type Truck struct {
	Car
	fourWheel bool
}

type Sedan struct {
	Car
	luxury bool
}

func main() {
	baseTruck := Car{
		doors: 5,
		color: "White",
	}
	baseSedan := Car{
		doors: 3,
		color: "Grey",
	}
	truck1 := Truck{
		Car:       baseTruck,
		fourWheel: true,
	}
	sedan1 := Sedan{
		Car:    baseSedan,
		luxury: false,
	}

	fmt.Println(truck1.Car.doors, truck1.Car.color, truck1.fourWheel)
	fmt.Println(sedan1.Car.doors, sedan1.Car.color, sedan1.luxury)
}
