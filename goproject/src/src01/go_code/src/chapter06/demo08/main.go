package main

import (
	"fmt"
)

func main(){
	var a = [] int {1,2,3,4,5,6,7,8,9,10,12}
	fmt.Println("开始删除前a",a)
	a = a[:len(a)-1] //删除尾部一个元素
	fmt.Println("删除尾部一个元素a",a)
	a = a[:len(a)-3] //删除尾部3个元素
	fmt.Println("删除尾部三个元素a",a)
	a = a[1:] //删除开头1个元素
	fmt.Println("删除开头1个元素a",a)
	a = a[2:] //删除开头2个元素
	fmt.Println("删除开头2个元素a",a)
	a = append(a[:0],a[1:]...) //删除开头1个元素
	fmt.Println("删除开头1个元素append*a",a)
	a = append(a[:0],a[2:]...)   //删除开头2个元素
	fmt.Println("删除开头2个元素append*a",a)
	var b = [] int {1,2,3,4,5,6,7,8,9,10,12,5}
	fmt.Println("开始删除前b",b)
	b = append(b[:2],b[2+1:]...) //删除中间一个元素
	fmt.Println("删除中间一个元素b",b)
	b = append(b[:2],b[2+2:]...) //删除中间两个元素
	fmt.Println("删除中间二个元素b",b)
	b = b[:1+copy(b[1:],b[1+1:])] //删除中间一个元素
	fmt.Println("删除中间一个元素b(copy)",b)
	b = b[:1+copy(b[1:],b[1+2:])] //删除中间两个元素
	fmt.Println("删除中间两个元素b(copy)",b)
	 
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo08\main.go
// 开始删除前a [1 2 3 4 5 6 7 8 9 10 12]
// 删除尾部一个元素a [1 2 3 4 5 6 7 8 9 10]
// 删除尾部三个元素a [1 2 3 4 5 6 7]
// 删除开头1个元素a [2 3 4 5 6 7]
// 删除开头2个元素a [4 5 6 7]
// 删除开头1个元素append*a [5 6 7]
// 删除开头2个元素append*a [7]
// 开始删除前b [1 2 3 4 5 6 7 8 9 10 12 5]
// 删除中间一个元素b [1 2 4 5 6 7 8 9 10 12 5]
// 删除中间二个元素b [1 2 6 7 8 9 10 12 5]
// 删除中间一个元素b(copy) [1 6 7 8 9 10 12 5]
// 删除中间两个元素b(copy) [1 8 9 10 12 5]