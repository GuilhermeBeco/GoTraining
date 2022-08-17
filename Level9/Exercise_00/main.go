package main

import (
	"fmt"
	
	"sync"
)
var wg sync.WaitGroup

func main() {	
	fmt.Println("Inside main")
	wg.Add(1)
	go firstRoutine()
	fmt.Println("After first routine")
	wg.Add(1)
	go secondRoutine()
	fmt.Println("After second routine")
	wg.Wait()
}

func firstRoutine(){
	fmt.Println("First Routine")
	wg.Done()
}

func secondRoutine(){
	fmt.Println("Second Routine")
	wg.Done()
}