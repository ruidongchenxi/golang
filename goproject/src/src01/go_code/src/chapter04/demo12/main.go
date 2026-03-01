package main

import (
	"fmt"
	//"math"
)


func main(){
	var x int 
	fmt.Println("输入月份")
	fmt.Scanln(&x)
	switch {
	case x >=3 && x <=5 :
		fmt.Println("春季")
	case x >=6 && x <=8 :
		fmt.Println("夏季")
	case x >=9 && x <=11 :
		fmt.Println("秋季")
	case x <=2 || x == 12 :
		fmt.Println("冬季")
	default:
		fmt.Println("输入错误")
	}

		
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo12\main.go
// 输入月份
// 7
// 夏季
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo12\main.go
// 输入月份
// 10
// 秋季
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo12\main.go
// 输入月份
// 1
// 冬季
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo12\main.go
// 输入月份
// 2
// 冬季