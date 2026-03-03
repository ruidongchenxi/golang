package main

import (
	"fmt"
	util "src01/go_code/src/chapter05/demo02/utils"// 引用包,取别名
)
func main() {
	var  s int 
	var  x int 
	var  c string
	fmt.Println("输入第一个数")
	fmt.Scanln(&s)
	fmt.Println("输入第二个数")
	fmt.Scanln(&x)
	fmt.Println("输入操作符")
	fmt.Scanln(&c)
	re := util.Cal(s,x,c) //调用包里的函数，用别名访问
	fmt.Println(re)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run .\src\chapter05\demo02\main\mian.go
// 输入第一个数
// 23
// 输入第二个数
// 87
// 输入操作符
// *
// 2001