package main

import "fmt"

type Square struct {
	sideLengh float32
}

type Circle struct {
	radious float32
}

type Shape interface{
	area() float32
}

func (s Square) area() float32{
	return s.sideLengh*s.sideLengh
}

func(c Circle) area() float32{
	return c.radious*c.radious*3.14
}

func printArea(s Shape){
	fmt.Println(s.area())
}




func main() {
	s1 := Square{
		sideLengh: 7.5,
	}
	c1:=Circle{
		radious: 3.45,
	}
	printArea(s1)
	printArea(c1)

}
