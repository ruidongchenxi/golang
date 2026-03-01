package main
import (
	"fmt"
	"math"
)
func main(){
// 	var c int
// 	fmt.Println("输入成绩")
// 	fmt.Scanln(&c)
// 	if c == 100{
// 		fmt.Println("奖励一辆bwm")
// 	}else if c > 80 {
// 		fmt.Println("奖励iPhone17")	
// 	}else if c >=60 {
// 		fmt.Println("奖励iPad")
// 	}else{
// 		fmt.Println("继续努力吧")
// 	}
// 	var b bool
// //	if b = false{// golang if 条件判断不支持赋值;syntax error: cannot use assignment b = false as value
// 	if b == false{
// 		fmt.Println("a")
// 	}else if b {
// 		fmt.Println("b")
// 	}else if !b {
// 		fmt.Println("c")
// 	}else{
// 		fmt.Println("d")
// 	}
	var a float64 = 3.0
	var w float64 = 100.0
	var i float64 = 6.0
	m := w * w - 4 * a * i
	if m > 0 {
		x1 := (-w + math.Sqrt(m))/2*a
		x2 := (-w - math.Sqrt(m))/2*a
		fmt.Printf("x1=%v x2=%v\n",x1,x2)
	}else if m == 0{
		x1 := (-w + math.Sqrt(m))/2*a
		fmt.Printf("x1=%v x2=%v\n",x1)
	} else {
		fmt.Println("无解")
	} 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// 输入成绩
// 100
// 奖励一辆bwm
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// 输入成绩
// 78
// 奖励iPad
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// 输入成绩
// 85
// 奖励iPhone17
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// 输入成绩
// 45
// 继续努力吧
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// # command-line-arguments
// src\chapter04\demo04\main.go:19:7: syntax error: cannot use assignment b = false as value
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo04\main.go
// 输入成绩
// 80
// 奖励iPad
//a