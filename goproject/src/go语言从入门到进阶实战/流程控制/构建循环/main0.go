package main
import (
	"fmt"
)
func main(){
	var i int
	for ;; i++{
		if i > 10{
			break
		}
		fmt.Println(i)
	}
	var e int
	for{
		if e > 10{
			break
		}
		fmt.Println(e)
		e++
	}
	var a int
	for a < 10 {
		fmt.Println(a)
		a++
	} 
}