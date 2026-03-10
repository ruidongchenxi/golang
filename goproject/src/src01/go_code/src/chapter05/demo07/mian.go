package main

import (
	"fmt"
)

func sun2 (n1,n2 *int) {
	*n1,*n2=*n2,*n1
	
}
func main(){	
	
	n6 := 12
	n7 := 56
	fmt.Printf("调用函数前n6=%v,n7=%v\n",n6,n7)
	sun2(&n6,&n7)
	fmt.Printf("调用函数后n6=%v,n7=%v\n",n6,n7)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo06\mian.go
// 调用函数前n6=12,n7=56
// 调用函数后n6=56,n7=12
