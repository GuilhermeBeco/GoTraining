package main

import (
	"fmt"
)

func main() {
	// Solution go routing
	c := make(chan int)

	go func(){
		c <- 42
	}()

	fmt.Println(<-c)
	
	//Solution buffered bi dimensional channel
	c1 := make(chan int, 1)

	c1 <- 55

	fmt.Println(<-c1)

}
