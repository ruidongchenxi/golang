package main
import (
	"fmt"
	//"math"
)
func main() {
	var cat int =1
	var  str string = "banana"
	fmt.Printf("%p %p\n",&cat,&str)
}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo16\main.go
// 0xc00000a0b8 0xc000026070