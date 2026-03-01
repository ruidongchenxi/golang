
package main


import (
	"fmt"
)

func main() {
	//var t int 
	// for i:=0;i<=6;i++{
	// 	for j :=6;j>=0;j--{
	// 		if (i+j) == 6 {
	// 		fmt.Printf("%d + %d = %d\n",i,j,(i+j) )
	// 		}
	// 	}
	// }
	// fmt.Println("*******************************")
//优化
	for i:=0;i<=1000;i++{
		j:= 1000-i
		fmt.Printf("%d + %d = %d\n",i,j,i+j )
	}	

}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo17\main.go
// 0 + 6 = 6
// 1 + 5 = 6
// 2 + 4 = 6
// 3 + 3 = 6
// 4 + 2 = 6
// 5 + 1 = 6
// 6 + 0 = 6
// *******************************
// 0 + 6 = 6
// 1 + 5 = 6
// 2 + 4 = 6
// 3 + 3 = 6
// 4 + 2 = 6
// 5 + 1 = 6
// 6 + 0 = 6