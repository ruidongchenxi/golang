package main
import (
	"fmt"
)
func main(){
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	case "cx":
		fmt.Println(3)
	default: //兜底
		fmt.Println(0)
	}
}


