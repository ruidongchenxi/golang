package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){
	//var i int
	//i = 10
	//交换两个值
	a := 9
	b := 2
	fmt.Println("a=",a,"b=",b)
	//临时变量
	t := a
	a = b
	b = t
	fmt.Println("a=",a,"b=",b)
	//复合型赋值的操作
	a +=17
	fmt.Println("a=",a)
	var c int 
	c = a+c
	fmt.Println(c)
	//赋值运算符的左边只能是变量
	var d int
	d = a 
	d = 8 + 8 *8
	//d = test() + 90
	d = 890 
	fmt.Println(d)

}



