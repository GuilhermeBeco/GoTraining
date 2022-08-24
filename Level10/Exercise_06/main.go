package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	c:=make(chan int)
	
	go orch(c)
	

	for v:=range(c){
		fmt.Println(v)
	}

}

func orch(c chan int){
	for i:=0; i<10; i++{
		wg.Add(1)
		go goroutine(c, i)
		
	}
	wg.Wait()
	close(c)
}

func goroutine(c chan <- int, m int){
	go func(){
		for i:=0; i<10; i++{
			c <- i*10+m
		}
		wg.Done()
	}()
	
}