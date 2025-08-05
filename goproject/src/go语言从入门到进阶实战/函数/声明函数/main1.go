package main
import (
	"fmt"
)
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}
func main() {
	a,b := namedRetValues()
	fmt.Println(a,b)
	fmt.Println(namedRetValues())
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\声明函数\main1.go
1 2
1 2
定义函数返回两个整型值，分别ab
调用函数
输出函数返回值
*/