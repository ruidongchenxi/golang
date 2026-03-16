package main

import (
	"fmt"
)

func main(){
	var t = [...][3]int {{1,2,3},{4,5,6}} 
	fmt.Println("使用for 遍历")
	for i:=0;i<len(t);i++{
		for x:=0;x<len(t[i]);x++{
			fmt.Printf("%v ",t[i][x])
		}
		fmt.Printf("\n t[%v]=%v\n",i,t[i])
	}
	fmt.Println("使用for-range循环")
	for i,v := range t{
		for s,x := range v{
			fmt.Printf("t的%v的数组下标%v的值是%v\n",i,s,x)
			
		}
		fmt.Printf("t[%v]=%v\n",i,v)
	}


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo06\main.go
// 使用for 遍历
// 1 2 3 
//  t[0]=[1 2 3]
// 4 5 6
//  t[1]=[4 5 6]
// 使用for-range循环
// t的0的数组下标0的值是1
// t的0的数组下标1的值是2
// t的0的数组下标2的值是3
// t[0]=[1 2 3]
// t的1的数组下标0的值是4
// t的1的数组下标1的值是5
// t的1的数组下标2的值是6
// t[1]=[4 5 6]