package main

import "fmt"



func main() {
	x := returnWhatsOnePlusOne()
	y := returnWhatsOnePlusOne()
	x()
	x()
	y()
	y()
	y()

}

func returnWhatsOnePlusOne() func(){
	x := 10
	return func(){
		x++
		fmt.Println(x)
	}
}
