package main

import "fmt"

func main()  {
	//交换两个变量的值
	var a  int = 10
	var b  int = 20
	var t int
	fmt.Println(a,b)
	t = a
	a = b
	b = t
	fmt.Println(a,b)
   cx := 6
   dx := 8
   fmt.Println(cx,dx)
   cx,dx = dx,cx//交换两个变量
   fmt.Println(cx,dx)



}
