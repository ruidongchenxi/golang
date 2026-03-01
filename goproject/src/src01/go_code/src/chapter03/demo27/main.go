package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){

	//位运算
	// fmt.Println(2&3)
	// fmt.Println(2|3)
	// fmt.Println(2^3)
	// fmt.Println(2^-2)
	//fmt.Println(2^3)
	//fmt.Println(2^3)

    a:=1>>2
	c:=1<<2
	fmt.Println("a=",a,"c=",c)




}
//执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo27\main.go
// a= 0 c= 4