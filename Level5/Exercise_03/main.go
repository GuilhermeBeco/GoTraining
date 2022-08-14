package main

import "fmt"

func main() {
	baseTruck := struct {
		doors int
		color string
	}{
		doors: 3,
		color: "Black",
	}

	fmt.Println(baseTruck.color, baseTruck.doors)

}
