package main
import (
	"fmt"
)
func main(){
	var age byte
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if  age > 18{
		fmt.Println("你年龄太大")
	} else {
		fmt.Println("你该上学了")
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo02\main.go
// 请输入年龄
// 21
// 你年龄太大
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo02\main.go
// 请输入年龄
// 13
// 你该上学了