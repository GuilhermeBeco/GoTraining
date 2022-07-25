package main
import "fmt"

func main () {
	x := [][]string{}
	x = append(x, []string{"james", "bond", "pastel"})
	x = append(x, []string{"miss", "moneypenny", "hello"})

	for _, v :=range(x){
		for _, t :=range(v){
			fmt.Println(t)
		}
	}
}
