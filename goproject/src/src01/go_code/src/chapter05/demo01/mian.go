package main
import (	
	"fmt"
)
//将计算功能放到一个函数内，需要使用，调用即可
func cal(s int,x int, c string)  (int){
	var d int 
	if c == "*" {
		// fmt.Println(s,c,x,"=",s*x)
		d =s*x
	} else if c == "/"{
		// fmt.Println(s,c,x,"=",s/x)
		d = s/x
	} else if c == "+"{
		// fmt.Println(s,c,x,"=",s+x)
		d = s+x
	}else if c == "-"{
		// fmt.Println(s,c,x,"=",s-x)
		d = s - x
	}else if c == "%" {
		//fmt.Println(s,c,x,"=",s%x)
		d = s % x
	}
	return d
}
func main() {
	var  s int 
	var  x int 
	var  c string
	fmt.Println("输入第一个数")
	fmt.Scanln(&s)
	fmt.Println("输入第二个数")
	fmt.Scanln(&x)
	fmt.Println("输入操作符")
	fmt.Scanln(&c)
	re := cal(s,x,c)
	fmt.Println(re)
	//fmt.Println(as)
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo01\mian.go
// 输入第一个数
// 6
// 输入第二个数
// 8
// 输入操作符
// *
// 48