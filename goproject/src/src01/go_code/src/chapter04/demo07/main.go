package main
import (
	"fmt"
	//"math"
)
func main(){
	var s int 
	var x bool
	fmt.Println("请输入成绩")
	fmt.Scanln(&s)
	
	if s < 8 {
		fmt.Println("请输入性别")
		fmt.Scanln(&x)
		if x {
			fmt.Println("男子队进入决赛")
		}else{
			fmt.Println("女子对进入决赛")
		}
		
	}else{
		fmt.Println("很遗憾未进入决赛")
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo07\main.go
// 请输入成绩
// 5
// 请输入性别
// true
// 男子队进入决赛
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo07\main.go
// 请输入成绩
// 5
// 请输入性别
// false
// 女子对进入决赛
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo07\main.go
// 请输入成绩
// 9
// 很遗憾未进入决赛