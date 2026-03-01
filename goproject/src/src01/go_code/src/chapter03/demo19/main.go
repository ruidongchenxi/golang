package main

import (
	"fmt"
)

 func main(){
	//演示关系运算符

	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1==n2)//false
	fmt.Println(n1!=n2)//true
	fmt.Println(n1>n2)//true
	fmt.Println(n1>=n2)//true
	fmt.Println(n1<n2)//false
	fmt.Println(n1<=n2)//false
	fmt.Println(0.5*0.4)
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