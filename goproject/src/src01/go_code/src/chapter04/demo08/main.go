package main

import (
	"fmt"
	//"math"
)
func main(){
	var s int 
	var y int 

	fmt.Println("输入当前月份")
	fmt.Scanln(&y)
	

   // if y == 4 || y == 5 || y == 6 || y == 7 || y == 8 || y == 9 || y == 10 {
	if  y > 1 && y < 12 {  
		fmt.Println("输入年龄")
		fmt.Scanln(&s)
	   if  y >=4 && y <=10{
			if s >= 18 && s <= 60 {
				fmt.Println("旺季成年人票价60")
			}else if s <18 {
				fmt.Println("旺季儿童票价半价30")
			}else {
				fmt.Println("旺季老年人票价1/3 20")
			}
		} else {
			if s >= 18 && s <= 60 {
				fmt.Println("淡季成年人票价40")
			}else{
				fmt.Println("淡季儿童与老人票价20")
			}
		}
	} else {
		fmt.Println("输入当前月份错误")
	}	 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo08\main.go
// 输入当前月份
// 14
// 输入当前月份错误
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo08\main.go
// 输入当前月份
// 11
// 输入年龄
// 15
// 淡季儿童与老人票价20
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo08\main.go
// 输入当前月份
// 6
// 输入年龄
// 19
// 旺季成年人票价60
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo08\main.go
// 输入当前月份
// 7
// 输入年龄
// 65
// 旺季老年人票价1/3 20