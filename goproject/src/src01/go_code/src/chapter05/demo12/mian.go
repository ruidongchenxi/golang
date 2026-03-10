package main

import (
	"fmt"
)
func sun(n1 int,n2 int) int{
	defer fmt.Println("ok1 n1=",n1)//先进后出
	defer fmt.Println("ok2 n2=",n2)// 后进先出
	n1 ++ //n1=3
	n2 ++ //n2=5
	res := n1+n2  //8
	fmt.Println("ok res=",res)
	return res
}
func main(){	

	sun(2,4)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo10\mian.go
// ok res= 8
// ok2 n2= 4
// ok1 n1= 2