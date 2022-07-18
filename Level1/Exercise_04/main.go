package main
import "fmt"

type Pastel int

var x Pastel

func main () {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}