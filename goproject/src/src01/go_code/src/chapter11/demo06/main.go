package main

import (
	"fmt"
	//v2 "math/rand/v2"
)
type Point struct{
	x int
	y int
}
func main(){
	var a interface{}
	var P Point= Point{1,3}
	a = P
	//var b Point
	//b,ok := a.(Point)//表示把a 转换成Point
	if b,ok:= a.(Point); ok {
		fmt.Println("b成功",b)
	}else{
		fmt.Println("类型不匹配转换失败",a)
	}
//	fmt.Println(b)

	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo06\main.go
// b成功 {1 3}
// {1 3}
	

