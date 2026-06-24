package main

import (
	"fmt"
	//"runtime/trace"
)
func main(){
	var a,b = 2,3
	var astr,bstr = "hello","bobby"
	fmt.Println(a+b)
	fmt.Println(astr+bstr)
	fmt.Println(3%2)
	var abool,bbool = true,false
	if abool&&bbool{
		fmt.Println("a")
	}
	if abool||bbool{
		fmt.Println("b")
	}
	if !abool{
		fmt.Println("s")
	}
	//位运算符，性能要求高考虑
	var A = 60
	var B = 13
	fmt.Println(A&B)
}