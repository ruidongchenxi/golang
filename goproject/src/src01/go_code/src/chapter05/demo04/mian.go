package main

import (
	"fmt"
	//"net"
)

// func test3 (n int) int {
// 	if (n==1 || n==2){
// 		return 1
// 	}else {
// 		return test3(n-1) + test3(n-2)
// 	}
// }
// func f(n int) int {
// 	if n == 1{
// 		return 3
// 	}else {
// 		return 2 *f(n-1)+1
// 	}
// }

func c(n  int) int{
	if n>10 || n<1{
		return 0
	}else if n==10{
		return 1
	}else{
		return (c(n+1)+1)*2
	}
}
func main() {
	// fmt.Println(test3(6))
	// fmt.Println(f(5))
	fmt.Println(c(1))

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo04\mian.go
// 1534