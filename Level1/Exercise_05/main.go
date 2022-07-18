package main
import "fmt"

type Pastel int

var x Pastel
var y int

func main () {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}