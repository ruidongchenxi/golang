
package main


import (
	"fmt"
)

func main() {
	x :=0
	j :=0
	for i:=1;i<100;i++{
		
		if i % 9 == 0 {
			j ++
			x += i
			//fmt.Printf("%d+%d=%d\n",x,i)
			
		}
	}
	fmt.Println("和",x)
	fmt.Println("数量",j)
	

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo16\main.go
// 和 594
// 数量 11
