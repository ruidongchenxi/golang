package main

import "fmt"
//全局变量
var t = "s"
var  (
	s = 8
	o = false
)

func main() {
	//go定义
	var a int = 1
	a1 := "string"
	var a2 string
	a2 = "s"
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	var w1,w2,w3 ="r",5,"t"
	fmt.Println(w1,w2,w3)
	/*
	注意：
	变量必须先定义后使用
	go 语言是静态语言，要求变量的类型和赋值类型一致
	局部变量名字不能冲突；局部变量可以和全局变量重名优先级局部变量高
	简洁变量定义必须是在函数内
	定义变量后未赋值都有默认值（零值）
	局部变量定义后一定要使用否则报错
	*/

}