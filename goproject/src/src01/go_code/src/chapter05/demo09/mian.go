package main

import (
	"fmt"
)
var (
	//Fun1 就是全局匿名函数
	Fun1 = func (n1 int,n2 int) int {
		return n1*n2
	}
)

func sun2 (n1,n2 *int) {
	*n1,*n2=*n2,*n1
	
}
func main(){	
	res :=func (n1 int,n2 int ) int {
		return n1+n2
	}(2,3)//直接定义并调用匿名函数
	fmt.Println("res=",res)
	sub := func (n1 int,n2 int) int {
		return n1-n2
	}//将匿名函数赋值给一个变量
	fmt.Println("sub",sub(9,4)) //通过变量调用
	res1 := Fun1(9,6)// 全局匿名函数的使用
	fmt.Println("res1=",res1)
}
// 执行结果
// res= 5
// sub 5
// res1= 54
