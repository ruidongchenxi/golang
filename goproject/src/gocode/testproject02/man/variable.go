package main
import "fmt"

func main(){
	//1.变量声明
	var age int
	//2.变量赋值
	age = 17
	//3.变量使用
	fmt.Println("age=" , age)
	//变量的重复定义错误
	//var age int = 9
	
	/*   重复定义报错不是重新赋值
	D:\golang\goproject\src\gocode\testproject02\man>go run variable.go
    D:\golang\goproject\src\gocode\testproject02\man>go run variable.go
    # commandlinearquments
variable.go:11:6:age redeclared in this block
variable.go:6:6:other declaration of age
	*/
    fmt.Println("age=",age)
}