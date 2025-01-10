package main
import (
	"fmt"
)
func main(){
	var ten int = 11
	if ten > 10 {
		fmt.Println(">10")
	}else if ten == 10 {
		fmt.Println("=10")
	}else{
		fmt.Println("<10")
	}
}