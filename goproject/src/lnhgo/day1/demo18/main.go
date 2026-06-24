package main

import (
	"fmt"
)
func swap (a,b *int){
	*a,*b=*b,*a
}
func main() {
    x, y:= 1,2
	fmt.Println(x,y)
	swap(&x,&y)
	fmt.Println(x,y)

}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo18\main.go
// 1 2
// 2 1