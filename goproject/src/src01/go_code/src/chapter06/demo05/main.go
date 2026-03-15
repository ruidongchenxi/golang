package main

import (
	//"bytes"
	"fmt"
	"math/rand"
)
func test (arr [3]int){
	arr[0] = 88
	fmt.Println("test",arr)
}
func test1 (arr *[3]int){
	(*arr)[0]=90 //（*arr[]）
	fmt.Println("test1",*arr)

}
func main(){
	var t [6] int //初始化数组
	for i,_ :=range t{
		t[i]=rand.Intn(100)+1 //数组下标赋值
	}
	fmt.Println("交换前",t) //赋值后打印
	y :=0
	for i :=len(t)-1;i>=len(t)/2;i--{
		t[i],t[y] = t[y],t[i] //交换两个值
		y++
	}
	fmt.Println("交换后",t) //交换后打印
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo05\main.go
// 交换前 [18 85 29 93 13 5]
// 交换后 [5 13 93 29 85 18]