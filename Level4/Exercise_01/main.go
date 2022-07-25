package main
import "fmt"

func main () {
	var x = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	fmt.Printf("%T\n", x)
	for k,v := range x{
		fmt.Println(k,v)
	}
}
