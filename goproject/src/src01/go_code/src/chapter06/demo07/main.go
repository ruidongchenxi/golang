package main

import (
	//"bytes"
	"fmt"
	_ "math/rand"
	_"net"
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
	// var c []int = []int {1,2,4}
	// fmt.Println("c=",c)
	// c = append(c[:1],append([]int{10},c[1:]...)...)//在下标1的位置插入10
	// fmt.Println("c=",c)
	// c = append(c[:1],append([]int{23,45,89},c[1:]...)...) //在下标1的位置插入切片
	// fmt.Println("c=",c)
	// c = append(c, 0) //切片扩展1个空间
	// copy (c[3+1:],c[3:])//c[3:]向后移动一个位置;整体向后移动一个位置 因为从第四个元素到后面所有元素拷贝。从第四（下标是3下标从0开始）个元素开始所有元素整体向后移，因为是拷贝下标是3的四个元素此时还是原来放入值；因此在下一步就是c[3]=8,就完成了添加元素
	// fmt.Println("copy操作后，c=",c)

	// c[3]=8// 重新设置下标为3的值
	// fmt.Println("copy操作并重新赋值后，c=",c)
	// d := c[1:4]
	// fmt.Println(d)
	// fmt.Println(&c[1])
	// fmt.Println(&d[0])
	// fmt.Println(cap(d))
	// // d = append([]int{0},d...)
	// d =append(d, 1,2,3)
	// fmt.Println(&c[1])
	// fmt.Println(&d[0])
	// var t [5]int = [5]int {3,7,8,8,5}

	// e := t[:5]
	// fmt.Println("数组1",&t[0])
	// fmt.Println("e1",&e[0])
	// e = append(e,1,2,3 )
	// fmt.Println("数组",&t[0])
	// fmt.Println("e",&e[0])
	var g []int = []int {1,2,3}
	fmt.Println(g)
	g = append(g, 4)
	fmt.Println(g)
	var b []int = []int {1,2,3,4,5}
	var c = make([]int,10)
	fmt.Println(c)
	copy(c,b)
	fmt.Println(c)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo07\main.go
// [1 2 3]
// [1 2 3 4]
// [0 0 0 0 0 0 0 0 0 0]
// [1 2 3 4 5 0 0 0 0 0]