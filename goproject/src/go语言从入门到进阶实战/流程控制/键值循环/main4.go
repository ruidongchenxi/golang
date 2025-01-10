package main
import (
	"fmt"
)
func main(){
	m := map[string]int{
		"hello": 100,
		"work": 200,
	}
	for _, value := range m{
		fmt.Println(value)
	}
}