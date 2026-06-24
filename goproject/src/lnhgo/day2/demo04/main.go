package main
import (
	"fmt"
)
func main(){
	var n int
	fmt.Println("输入：")
	fmt.Scanln(&n)
	switch n{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("输入有误")
	}
}
/*
PS D:\golang\goproject\src\lnhgo> go run day2\demo04\main.go
输入：
7
输入有误
PS D:\golang\goproject\src\lnhgo> go run day2\demo04\main.go
输入：
3
中指
*/