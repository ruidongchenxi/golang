package main
import (
	"fmt"
)
const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo06\main.go
// 1
// 2
// 2
// 3
// 3
// 4