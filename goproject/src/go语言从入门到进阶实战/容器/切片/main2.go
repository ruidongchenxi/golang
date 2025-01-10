package main
import (
	"fmt"
)
func main(){
	a := [] int {1,2,3}
	fmt.Println(a[:])
}
/*
执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main2.go
[1 2 3]
a 是一个拥有3个元素的切片，将a 使用a[:]进行操作后，得到的切片与a切片一致
*/