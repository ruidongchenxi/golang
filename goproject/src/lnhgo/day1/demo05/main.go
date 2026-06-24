package main
import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	)
func main(){
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo05\main.go
// 1024
// 1048576
// 1073741824
// 1099511627776
// 1125899906842624