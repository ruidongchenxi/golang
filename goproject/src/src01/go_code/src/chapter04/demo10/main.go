package main

import (
	"fmt"
	//"math"
)


func main(){


   var i int = 10 
   switch i {
   case 10:
		fmt.Println("10")
		fallthrough // 穿透
   case 11:
		fmt.Println("12") //执行
   case 12:
		fmt.Println("13")
   } 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo09\main.go
// 11
// 12
