package main
import "fmt"
func main(){
	fmt.Println(add(30,60))

}
func add(num1 int,num2 int) int{
	//在golang中，程序遇到defer关键字不会立即执行defer后的语句，而是将defer后的语句压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1 =",num1)
	defer fmt.Println("num2 =",num2)
	num1 += 90
	num2 += 30
	fmt.Println(num1,num2)
	//栈的特点先进后出
	//函数执行完成，从栈中取出语句开始执行
	var sum int = num1 + num2
	fmt.Println("sum=",sum)
	return sum
}