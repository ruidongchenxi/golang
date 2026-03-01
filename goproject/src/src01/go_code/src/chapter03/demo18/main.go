package main

import (
	"fmt"
	_"strconv"
)

 func main(){
	//在Golang中，++和--只能独立
	var i int =8
	var a int
	i++
	a = i
	fmt.Println("a=",a)
	if i > 0 {
		fmt.Println("ok")
	}
}
// 执行结果
// 2
// n1 = 2
// n2 = 2.5
// 10%3 1
// -10%3 -1
// 10%-3 1
// -10%-3 -1
// i= 11
// i= 10