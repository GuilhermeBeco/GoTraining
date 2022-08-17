package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementer int
var wg sync.WaitGroup
var mu sync.Mutex

func main(){
	for i:=0; i<100; i++{
		wg.Add(1)
		go routine()
	}
	wg.Wait()
	fmt.Println(incrementer)
}

func routine(){
	mu.Lock()
	defer mu.Unlock()
	
	localIncremeneter :=  incrementer
	runtime.Gosched()
	localIncremeneter++
	incrementer=localIncremeneter
	wg.Done()
}