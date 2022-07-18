package main
import "fmt"

func main () {
	fmt.Printf("true && true = %t\n", true && true)
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("true || false = %t\n", true || false)
	fmt.Printf("true || true = %t\n", true || true)
	fmt.Printf("!true = %t\n", !true)
}

