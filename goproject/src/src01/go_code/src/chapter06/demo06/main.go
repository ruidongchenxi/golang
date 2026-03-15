package main

import (
	//"bytes"
	"fmt"
	_ "math/rand"
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
	var(
		// a []int //nil切片，和nil相等，表示一个不存在的切片
		// b = []int {} // 空切片。和nil不相等，一般用来表示一个空集合
		//c = []int{1,2,3}//有3个元素的切片，len为2，cap为3
		// d = c[:2] //有两个元素的切片，len为2，cap为3；c开头到小标为2 
		// e = c[0:2:cap(c)] //有2个元素的切片。len为2，cap为3
		// f = c[:0] //有0个元素，len为0，cap为3
		// g = make([]int,3) //有3个元素的切片。len和cap都为3；cap 省略，表示与
		// h = make([]int,2,3) //有2个元素的切片。len为2，cap为3
		// i = make([]int,0,3)//有0个元素的切片。len为0，cap为3
	)
	// var in [5]int = [5] int {1,2,3,9,7}
	// s := in[1:3] //s：是切片的名字in[1:3]表示从下标为1的元素开始到下标为3的元素前一个，包左不包右
	// fmt.Println("数组",in)
	// fmt.Println("切片=",s)
	// fmt.Println("切片 长度=",len(s))
	// fmt.Println("切片 容量=",cap(s))//切片的容量是可以动态变化的；cap = 原数组长度 - 起始位置
	// fmt.Println("in len=",len(in))
	// fmt.Printf("in[1]的地址值%p\n",&in[1]) //打印数组下标为1的元素内存地址值
	// fmt.Printf("s[0]的地址值%p\n",&s[0]) //打印切片的内存地址值
	// fmt.Println("in 数组",in)
	// s[0]= 4 //修改切片第0个元素
	// fmt.Println("in 数组",in)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo06\main.go
// in [1 2 3 9 7]
// slice= [2 3]
// slice len= 2
// slice cap= 4
// in len= 5
// in[1]的地址值0xc00000e338
// s[0]的地址值0xc00000e338
// in 数组 [1 2 3 9 7]
// in 数组 [1 4 3 9 7]
