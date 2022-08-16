package main

import "fmt"



func main() {
	x := returnWhatsOnePlusOne()
	fmt.Println(x())
}

func returnWhatsOnePlusOne() func()int{
	return func() int{
		return 1+1
	}
}
