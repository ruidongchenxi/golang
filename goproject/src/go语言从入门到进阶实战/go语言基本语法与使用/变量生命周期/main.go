package main
import (
	"fmt"
)
//本函数测试入口参数和返回值情况
func dummy(b int) int{
	//本函数测试入口参数和返回值情况
	//声明一个c赋值进入参数并返回
	var c int
	c = b
	return c
}
//空函数，什么也不做
func void(){
	}
func main(){
	//声明变量a并打印
	var a int
	//调用void()函数
	void()
	//打印a变量的值和dummy()函数
	fmt.Println(a,dummy(0))
}