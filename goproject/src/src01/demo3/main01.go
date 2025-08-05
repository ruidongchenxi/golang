package main

import (
	"fmt"
	"unsafe"
)
func main(){
	//定义一个整型
	var num1 int8 = 23

	fmt.Println(num1)
	/*
		D:\golang>go run D:\golang\goproject\src\src01\demo3\main01.go
		# command-line-arguments
		goproject\src\src01\demo3\main01.go:5:18: cannot use 230 (untyped int constant) as int8 value in variable declaration (overflows)

	*/
	var num3 = 28
	//printf 函数作用就是格式化
	/*
	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%T	值的类型的Go语法表示
	%%	百分号
	 */
	fmt.Printf("num3的类型：%T", num3)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3))//unsafe 导包。显示字节数

}