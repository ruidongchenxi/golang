package main

import (
	"fmt"
	) 

var name = "tom"
//Name := "cx" 等价var Name string Name="cx" 因为初始化可以，赋值属于执行操作，必须在函数里执行；src\chapter05\demo14\mian.go:8:1: syntax error: non-declaration statement outside function body

func test01() {
	fmt.Println("全局变量",name)
}
func test02(){
	name := "jack"
	fmt.Println("test02局部变量",name)
}
func main(){
	fmt.Println("main函数",name)
	test01()
	test02()
	test01()
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo14\mian.go
// # command-line-arguments
// src\chapter05\demo14\mian.go:8:1: syntax error: non-declaration statement outside function body