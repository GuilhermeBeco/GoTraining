package main

import "fmt"



func main() {
	x := func()int{
		return 1+1
	}()
	fmt.Println(x)
}
