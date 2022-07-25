package main
import "fmt"

func main () {
	var x = [5]int{1,2,3,5,5}
	fmt.Printf("%T\n", x)
	for k,v := range x{
		fmt.Println(k,v)
	}
}
