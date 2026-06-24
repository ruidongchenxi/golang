package main
import (
	"fmt"
)

func main(){
	var i1 = 101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%o\n",i1)//10进制转8进制
	fmt.Printf("%x\n",i1)//10进制转16进制
	fmt.Printf("%b\n",i1)//10进制转2进制
	//八进制
	i2:= 077
	fmt.Printf("%d\n",i2)
	//十六进制
	i3:= 0xab3
	fmt.Printf("%d\n",i3)
	//查看变量类型
	fmt.Printf("%T\n",i3)

}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo07\main.go
// 101
// 145
// 65
// 1100101
// 63
// 2739
// int