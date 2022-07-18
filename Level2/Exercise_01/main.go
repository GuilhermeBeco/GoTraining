package main
import "fmt"

func main () {
	x:=	42
	y:= 74
	resEquals := x == y
	resSmallerOrEquals := x <= y
	resBiguerOrEquals := x >= y
	resDiferent := x != y
	resSmaller := x < y
	resBiguer := x > y
	fmt.Println(resEquals, resSmallerOrEquals, resBiguerOrEquals, resDiferent, resSmaller, resBiguer)


}
