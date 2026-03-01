
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("你好，晨曦")
	}
	fmt.Println("*********************************")
	j := 0
	for j < 3 {
		j++
		fmt.Println("你好，Thanks♪(･ω･)ﾉ", j)
	}
	fmt.Println("**********************************")
	y := 0
	for {
		//y := 0
		y++
		
		if y > 3 {//
			break
		}
		fmt.Println("你好", y)
	}
	fmt.Println("***************************************")
	t := 0
	for {
		if t < 3{
			t++
			fmt.Println("你好hh",t)
		}else{
			break
		}
		
	}
}


// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo14\main.go
// 你好，晨曦
// 你好，晨曦
// 你好，晨曦
// *********************************
// 你好，Thanks♪(･ω･)ﾉ 1
// 你好，Thanks♪(･ω･)ﾉ 2
// 你好，Thanks♪(･ω･)ﾉ 3
// **********************************
// 你好 1
// 你好 2
// 你好 3
// ***************************************
// 你好hh 1
// 你好hh 2
// 你好hh 3
