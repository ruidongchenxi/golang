package main
import (
	"fmt"
)
func calc(a,b int) int{
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}
func main(){
	var (
		cx int
		d int
	)
	cx = 9
	d = 5
	z := calc(cx,d)
	fmt.Println(z)
}