package main

import "fmt"



func main() {
	x := returnWhatsOnePlusOne()
	fmt.Println(x())
	addOneToOnePlusOne(x)
}

func returnWhatsOnePlusOne() func()int{
	return func() int{
		return 1+1
	}
}

func addOneToOnePlusOne(t func()int){
	fmt.Println(t()+1)
}
