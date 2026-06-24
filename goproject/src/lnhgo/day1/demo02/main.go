package main
import (
	"fmt"
)
// 常量
const pi = 3.1415926
//批量声明常量
const (
	n1 = 200
	n2 = 300
	n3   //n3后面没有写值默认跟上面一样

)
func main(){
	fmt.Println("n1:",n1)
	fmt.Println("n1:",n2)
	fmt.Println("n1:",n3)
	fmt.Println("pi:",pi)
}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo02\main.go
// n1: 200
// n1: 300
// n1: 300
// pi: 3.1415926