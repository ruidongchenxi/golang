package main
import (
	"fmt"
)
func main(){
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day2\demo01\main.go
// a
// b