package main
import (
	"fmt"
)
func main(){
	var c [3]int = [3]int {1,2,3}
	fmt.Println(c)
	var d = [3]int {1,2,4}
	fmt.Println(d)
	e := [...]int {4:5}
	fmt.Println(e)
	for i,v := range e{
		fmt.Printf("数组下标=%v,数组下标对应的值=%v\n",i,v)
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo02\main.go
// [1 2 3]
// [1 2 4]
// [0 0 0 0 5]
// 数组下标=0,数组下标对应的值=0
// 数组下标=1,数组下标对应的值=0
// 数组下标=2,数组下标对应的值=0
// 数组下标=3,数组下标对应的值=0
// 数组下标=4,数组下标对应的值=5