package main

import (
	"fmt"
	_"strconv"
)

 func main(){
	//指针
	var i int =10
	//i的地址是什么，&i
	var prt *int = &i
	fmt.Println(prt)
	//打印地址值变量的变量地址
	fmt.Println(&prt)
	//根据指针的变量打印指向的值
	fmt.Println("i =",i)
	*prt = 90
	fmt.Println("*prt=",*prt)
	fmt.Println("i=",i)

}
// 执行结果

// 0xc00000a0a8
// 0xc000058028
// i = 10
// *prt= 90
// i= 90