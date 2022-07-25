package main
import "fmt"

func main () {
	var x = []int{1,2,3,4,5,6,7,8,9,10,11}
	y := x[1:6]
	fmt.Println(y)
	z := x[6:]
	fmt.Println(z)
	w:= x[3:9]
	fmt.Println(w)
	p:=x[2:8]
	fmt.Println(p)
}
