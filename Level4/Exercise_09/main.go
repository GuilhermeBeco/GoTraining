package main
import "fmt"

func main () {
	map1 := map[string][]string{
	"Last1_First1":[]string{"one","two","three"},
	"Last2_First2":[]string{"four","five","six"},
	"Las32_First3":[]string{"seven","eigth","nine"},
	}
	map1["Last4_First4"] = []string{"ten","eleven","twelve"}

	delete(map1, "Last2_First2")
	for _, v:= range map1{
		fmt.Println(v)
	}
}
