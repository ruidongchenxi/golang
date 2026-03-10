package main

import (
	"fmt"
	) 

func s (t int) {
	for i :=1;i <=t;i++{
        for k :=1; k<=t-i;k++{
            fmt.Printf(" ")
        }
        for g :=1; g<=2*i-1;g++{
            fmt.Printf("*")
        }
        fmt.Println()
    }
}
func x(r int) {
	for i:=1;i<=r;i++{
        //fmt.Println(i)
        for j := 1; j <= i ;j++{
            fmt.Printf("%d * %d = %d ",j,i,(i*j))
            fmt.Printf("/t")
        }
        fmt.Println("")
    }
}
func main(){
	//fmt.Println("main函数",name)
	s(8)
	x(15)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo15\mian.go
//        *
//       ***
//      *****
//     *******
//    *********
//   ***********
//  *************
// ***************
// 1 * 1 = 1
// 1 * 2 = 2       2 * 2 = 4
// 1 * 3 = 3       2 * 3 = 6       3 * 3 = 9
// 1 * 4 = 4       2 * 4 = 8       3 * 4 = 12       4 * 4 = 16
// 1 * 5 = 5       2 * 5 = 10       3 * 5 = 15       4 * 5 = 20       5 * 5 = 25
// 1 * 6 = 6       2 * 6 = 12       3 * 6 = 18       4 * 6 = 24       5 * 6 = 30       6 * 6 = 36
// 1 * 7 = 7       2 * 7 = 14       3 * 7 = 21       4 * 7 = 28       5 * 7 = 35       6 * 7 = 42       7 * 7 = 49