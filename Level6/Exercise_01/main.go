package main

import "fmt"



func main() {
	x:= []int{1,2,3,4,5,6,7,8,9,10}
	totalFoo:=foo(x...)
	totalBar:=bar(x)
	fmt.Println(totalFoo)
	fmt.Println(totalBar)
}

func foo(ti ...int) (int) {
	total:=0
	for _, value :=range(ti){
		total+=value
	}
	return total
}
func bar(ti []int) (int) {
	total:=0
	for _, value :=range(ti){
		total+=value
	}
	return total
	
}
