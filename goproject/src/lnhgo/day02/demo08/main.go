package main

import (
	"fmt"
	//v2 "math/rand/v2"
	//"weak"
)

//
func add(a,b int,)int{
	//fmt.Println(c)
	return  a+b
}
//多个返回值
func add1(a,b int)(sun int,err error){
	sun=a+b
	return sun,err
}
//可变参数示例
func add2(a ...int)(sun int){
	for _,v:=range a{
		sun+=v
	}
	return
}
//前面这个string类型参数必须传，后面可变int类型参数可以不传
func add3(d string,c ...int)(sun int){
	for _,v:=range c{
		sun+=v
	}
	fmt.Println(d)
	return
}
func runForever(){
	for {

	}
}
//函数返回函数
func cal(op string,i ...int) func(){
	switch op{
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	case "*":
		return func() {
			fmt.Println("这是乘法")
		}
	case "/":
		return func() {
			fmt.Println("这是除法")
		}
	case "%":
		return func() {
			fmt.Println("这是取余")
		}
	default:
		return  func() {
			fmt.Println("输入符号不正确")
		}
	}
}
//传递的参数也带有函数
func cal2(my func(i ...int) int ) int{
	return my()
}
//闭包
func auo() func() int{
	local:=0//
	return func () int{
		local+=1
		return local
	}
}

func main(){
	q:=2
	s:=7
	t:=add(q,s)
	fmt.Println(t)
	//add函数赋值给a变量；注意不可带括号，带括号是调用
	a:=add
	//调用a变量函数返回值给c
	c:=a(5,8)
	fmt.Println(c)
	//调用
	cal2(func (i ...int) int{
		sun	:= 0
		for _,v :=range i{
			sun+=v
		}
		return sun
	})
}
