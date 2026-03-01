package main

import (
	"fmt"
	_"strconv"
)

 func main(){
	//重点讲解/ %
	//说明，如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10/4)
	var n1 float32 = 10 / 4
	fmt.Println("n1 =",n1)
	// 希望小数保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println("n2 =", n2)
	//%
	//公式a%b=a-a/b*b
	fmt.Println("10%3",10%3)
	fmt.Println("-10%3",-10%3)//-1
	fmt.Println("10%-3",10%-3)//1
	fmt.Println("-10%-3",-10%-3) //-1
	var i int = 10
	i ++ 
	fmt.Println("i=",i)
	i -- 
	fmt.Println("i=",i)

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