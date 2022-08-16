package main

import "fmt"



func main() {
	x:= foo()
	z,y := bar()

	fmt.Println(x)
	fmt.Println(z, y)
}

func foo() (int) {
	return 10
}
func bar() (int,string) {
	return 20, "HI"
}
