package main

import (
	//"crypto/x509"
	"fmt"
)
func main() {
	// for i:=1;i<10;i++{
	// 	if i % 2 == 0{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	var z int
	var f int 
	for {
		var x int 
		fmt.Println("输入一个数字")
		fmt.Scanln(&x)
		if x > 0{
			z++
		}else if x == 0{
			fmt.Println("输入为0无法继续输入")
			break
		}else{
			f++
		}
	}
	fmt.Printf("输入正数个数%v，负数个数%v",z,f)	
	

}
// 执行结果

// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo24\main.go
// 输入一个数字
// 89
// 输入一个数字
// 65
// 输入一个数字
// -98
// 输入一个数字
// 65
// 输入一个数字
// -97
// 输入一个数字
// -8
// 输入一个数字
// 3
// 输入一个数字
// 0
// 输入为0无法继续输入
// 输入正数个数4，负数个数3