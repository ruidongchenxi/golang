package main

import (
	"fmt"
	//"net"
)

func c(n  int) int{
	if n>10 || n<1{
		return 0
	}else if n==10{
		return 1
	}else{
		return (c(n+1)+1)*2
	}
}
func test() {
	//n1 是test 函数的局部变量，只能在test 函数中使用
	var n1 = 100
	fmt.Println(n1)
}
func test2(n1 *int) { //传指针(值传递)；n1就是*int
	//n1 是test 函数的局部变量，只能在test 函数中使用
	*n1+= 10
	fmt.Println("test2(),n1=",*n1)
}
func main() {
	// fmt.Println(test3(6))
	// fmt.Println(f(5))
	//fmt.Println(c(1))
	var n1 int = 16 
	fmt.Println("man() 调用test之前 n1=",n1)
	test2(&n1)//传指针(值传递)
	fmt.Println("man()调用test后 n1=",n1)
}
//执行结果
// man() 调用test之前 n1= 16
// test2(),n1= 26
// man()调用test后 n1= 26