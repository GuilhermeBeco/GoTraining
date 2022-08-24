package main

import (
	"fmt"
	"github.com/GuilhermeBeco/GoTraining/Level12/Exercise_00/dog"
)

func main(){
	humanYears := 23
	fmt.Println("Human years:", humanYears, "Dog years:",dog.Years(humanYears))
}