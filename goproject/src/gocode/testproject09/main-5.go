package main
import "fmt"
func test (num *int){
	*num = 30
	fmt.Println("test地址值-----------",num)
	fmt.Println("test值-----------",*num)

} 
func main(){
	var num int = 10
	test(&num)
	fmt.Println("main--------",num)
}