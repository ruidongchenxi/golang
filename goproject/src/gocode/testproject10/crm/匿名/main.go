package main
import "fmt"
var Func01 = func (num1 int, num2 int) int {
	return num1 * num2
}
func main(){
	result := func (num1 int, num2 int) int{
		return num1 + num2
	}(10,20)
	fmt.Println(result)
	var sub = func (num1 int,num2 int) int{
		return num1 - num2
	}
	chenxi := sub(79,45)
	fmt.Println(chenxi)
	fmt.Println(Func01(3,7))
}