package main

import (
	"fmt"
	//"math"
)


func main(){
	var s int
	fmt.Println("输入成绩")
	fmt.Scanln(&s)
	switch {
	case s > 100 :
		fmt.Println("输入错误")
	case s >= 60 :
		fmt.Println("合格")
	default:
		fmt.Println("不合格")
	}

		
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo11\main.go
// 输入成绩
// 80
// 合格