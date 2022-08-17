package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var incrementer int64
var wg sync.WaitGroup


func main(){
	for i:=0; i<100; i++{
		wg.Add(1)
		go routine()
	}
	wg.Wait()
	fmt.Println(incrementer)
}

func routine(){	
	atomic.AddInt64(&incrementer,1)
	fmt.Println(atomic.LoadInt64(&incrementer))
	wg.Done()
}