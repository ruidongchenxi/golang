package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){
	
	a := 100
	fmt.Println("a的地址值",&a)
	var ptr *int = &a
	fmt.Println("ptr 指向的值是=", *ptr)
	var n int
	var i int  = 10
	var j int  = 12
	if i > j {
		n= i
	}else{
		n = j
	}
	fmt.Println(n)

}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo22\main.go
// a的地址值 0xc00000a0a8
// ptr 指向的值是= 100
// 12


