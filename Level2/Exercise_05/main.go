package main
import "fmt"

const(
	x=2018 + iota
	y
	z
	w
)


func main () {
	fmt.Println(x, y, z, w)
}
