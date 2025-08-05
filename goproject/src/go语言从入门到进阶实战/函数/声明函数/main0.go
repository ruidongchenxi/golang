package main
import (
	"fmt"
)
func typedTwoValues() (int, int) {
	return 1,2
}
func main() {
	a, b := typedTwoValues()
	fmt.Println(a,b)
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\声明函数\main0.go
1 2
纯类型的返回值对于代码可读性不是很友好，特别是在同类型的返回值出现时，无法区分每个返回参数的意义
*/