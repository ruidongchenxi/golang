package main
import "fmt"
func fire() {
	fmt.Println("fire")
}
func main() {
	var f func()
	f = fire
	f()
}
/*
定义一个fire()函数
将变量f声明为func()类型，此时f就被俗称为“回调函数”，此时f的值为nil
将fire()函数名作为值赋给f变量 
使用f变量进行函数调用
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\函数变量\main0.go
fire
*/