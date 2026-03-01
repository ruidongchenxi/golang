package main

import (
	"fmt"
	//"math"
)


func main(){
	var q string
	fmt.Println("输入对应星期几")
	fmt.Scanln(&q)
	switch q{
	case "星期一":
		fmt.Printf("干煸豆角")
	case "星期二":
		fmt.Printf("醋溜土豆")
	case "星期三":
		fmt.Printf("红烧狮子头")
	case "星期四":
		fmt.Printf("油炸花生米")
	case "星期五":
		fmt.Printf("蒜蓉扇贝")
	case "星期六":
		fmt.Printf("东北乱炖")
	case "星期日":
		fmt.Printf("大盘鸡")
	default:
		fmt.Println("输入错误")
	}
		
}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo13\main.go
// 输入对应星期几
// 星期一
// 干煸豆角