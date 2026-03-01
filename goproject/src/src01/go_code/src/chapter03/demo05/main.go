package main

import (
	"fmt"
	"unsafe"
)
func main(){

	var n1 = 120 
	fmt.Printf("n1 的类型%T\n",n1)
	var n2 int64 = 10
	//unsafe.Sizeof()是一个unsafe包里的函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的类型%T，n2占用的字节数是%d\n",n2,unsafe.Sizeof(n2))
	var age byte =90
	fmt.Println(age)

}
// 执行结果
// n1 的类型int
// n2 的类型int64，n2占用的字节数是8