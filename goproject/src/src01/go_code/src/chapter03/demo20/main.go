package main

import (
	"fmt"
)
// 声明一个函数
func test() bool{
	fmt.Println("test........")
	return true
}
 func main(){
	//逻辑运算符的使用
	var i int = 10
	//短路与
	if i < 9 && test(){
		fmt.Println("ok....")
	} 
	//短路或
	if i < 9 || test(){
		fmt.Println("hello....")
	} 
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo20\main.go
// test........
// hello....