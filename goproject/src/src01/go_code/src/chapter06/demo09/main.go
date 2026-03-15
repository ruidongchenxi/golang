package main

import (
	"fmt"
	//"slices"
)

func main(){
	str := "hello@chenxi.com"
	slice := str[6:]
	for _,v := range slice{
		fmt.Printf("%q",v)
	} 

	fmt.Println("\n",slice)
	

	//修改string
	//slice2 := str[:]
	fmt.Println("打印初始数组",str)
	arr := []byte(str) //转为切片
	arr[0] = 'z' //修改下标为0 的元素
	str = string(arr) //重新转成string 并赋值给原来变量
	fmt.Println("打印修改后数组",str)
	//细节，我们转成byte后，可以处理英文和数字，无法处理中文;原因是byte是按字节编码的
	//解决方法：[]rune 是按字符处理
	arr1 := []rune(str)
	arr1[0]= '陈'
	str = string(arr1)
	fmt.Println("rune 修改汉字",str)

	
	



}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo09\main.go
// 'c''h''e''n''x''i''.''c''o''m'
//  chenxi.com
// 打印初始数组 hello@chenxi.com
// 打印修改后数组 zello@chenxi.com
// 陈ello@chenxi.com