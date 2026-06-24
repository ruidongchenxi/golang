package main
import (
	"fmt"
)

func main() {
	f1 := 1.23456
	fmt.Printf("%T\n",f1)//默认
	f2 := float32(1.234)//显示声明
	fmt.Printf("%T\n",f2)
	//f1 =
}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo10\main.go
// float64
// float32