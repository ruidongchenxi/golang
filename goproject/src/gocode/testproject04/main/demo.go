package main
import (
	"fmt"
    "unsafe"
)	

func main(){
	//var num1 int8 = 23
	//var num2 int8 =239 超出类型范围
	/*
	D:\golang\goproject\src\gocode\testproject03\man>go run demo.go                                                         
# command-line-arguments                                                                                                
.\demo.go:5:17: cannot use 239 (untyped int constant) as int8 value in variable declaration (overflows)                 
	*/
	//Printf 函数
	var num2 = 235
	fmt.Printf("num2的类型是:%T",num2)
	fmt.Println() //换行
	fmt.Println(unsafe.Sizeof(num2))
	//字符类型： 
}