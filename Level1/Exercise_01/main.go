package main
import "fmt"

var x int
var y string
var z bool

func main () {
	
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println("------")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
