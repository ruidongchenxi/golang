package main
import "fmt"
func main(){
    var age int = 18
	fmt.Println(&age)
	var s1 *int = &age
	fmt.Println(*s1)
} 