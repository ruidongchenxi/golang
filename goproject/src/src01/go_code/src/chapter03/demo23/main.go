package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){
	//基本赋值
	// var i int
	// i = 10 //基本赋值
	// 交换两个变量 
	var a int = 10
	var b int = 20
	fmt.Println(a)
	fmt.Println(b)
	// a = a + b
	// b = a - b
	// a = a - b
	a,b = b,a//交换两个变量的值
	
	fmt.Println(a)
	fmt.Println(b)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo23\main.go
// 10
// 20
// 20
// 10
