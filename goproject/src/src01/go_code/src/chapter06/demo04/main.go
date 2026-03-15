package main

import (
	//"bytes"
	"fmt"
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
	var mychars [26]byte
	for i,_ := range mychars{
		mychars[i] = 'A' + byte(i)
	}
	for i,_:= range mychars{
		fmt.Printf("%c ",mychars[i])
	}
	fmt.Println()
	fmt.Println("------------")
	//
	var arr1  = [9]int  {9,86,75,56,156,785,567,877,678}
	var x int
	var t int 
	for i,_ := range arr1{
		if arr1[i] >= x {
			x =arr1[i]
			t = i
		}


	}
	fmt.Printf("最大值=%v;最大值对应的下标是=%v\n",arr1[t],t)
	//求平均值跟和
	var arr2 [6]int = [6]int {56,78,98,45,56,89}
	var s int
	//var g int 
	for _,v := range arr2{
		s += v
	}
	//如何让平均值保留小数
	fmt.Printf("arr2 数组的和=%v，数组的平均值是=%v\n",s,float64(s)/float64(len(arr2)))
	
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo04\main.go
// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
// ------------
// 最大值=877;最大值对应的下标是=7
// arr2 数组的和=422，数组的平均值是=70.33333333333333