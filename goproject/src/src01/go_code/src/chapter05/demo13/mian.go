package main

import (
	"fmt"
	//"image"
)
func sun(n1 int,n2 int) int{
	defer fmt.Println("ok1 n1=",n1)//先进后出
	defer fmt.Println("ok2 n2=",n2)// 后进先出
	n1 ++ //n1=3
	n2 ++ //n2=5
	res := n1+n2  //8
	fmt.Println("ok res=",res)
	return res
}
//函数外定义的变量；称为全局变量 
var age int = 50 // 首字母小写整个包可以使用
var Name string = "cx" //首字母大写整个程序可以使用
//局部变量
func test () {
	//name和age变量的作用域，仅限于test()函数内部
	Name:="tom"
	age:=10
	fmt.Println(Name,age)

}
func main(){	

	fmt.Println("age=",age)
	if true{
		i := 10
		fmt.Println(i)
	}
	i:=89
	fmt.Println("if 里i",i)//i 在if 代码块里定义，所以生效范围仅限于if代码块里； src\chapter05\demo13\mian.go:34:24: undefined: i
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo13\mian.go
// # command-line-arguments
// src\chapter05\demo13\mian.go:34:24: undefined: i