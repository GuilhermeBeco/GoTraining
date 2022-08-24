package main

import(
	"fmt"
)
type customErr struct{
	info string
}

func (ce customErr) Error() string{
	return fmt.Sprintf("OH NO AN ERROR: %v", ce.info)
}

func main(){
	ce := customErr{
		info:"pastel",
	}

	foo(ce)
}

func foo(e error){
	fmt.Println("foo ran -", e, "\n")
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}