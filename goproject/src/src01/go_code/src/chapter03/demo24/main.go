package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){
	//var y int
	var q int = 5
	var w int = 9
	fmt.Println("两个数最大值")
	if q > w {
		fmt.Println("最大值：",q)
	}else{
		fmt.Println("最大值：",w)
	}
	fmt.Println("三个数最大值")
	d := 8
	c := 4
	f := 9
	var y int 
	if d > c {
		y = d
	}else{
		y = c
	}
	if f >y{
		y = f
	}
	fmt.Println("最大值",y)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo24\main.go
// 两个数最大值
// 最大值： 9
// 三个数最大值
// 最大值 9
