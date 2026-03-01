
package main

import (
	"fmt"
)

func main() {
	str := "hello,world!北京"
	//str = []rune(str) //
	for i,v := range str{
		fmt.Printf("index=%d val=%c\n",i,v)
	}
	

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo15\main.go
// index=0 val=h
// index=1 val=e
// index=2 val=l
// index=3 val=l
// index=4 val=o
// index=5 val=,
// index=6 val=w
// index=7 val=o
// index=8 val=r
// index=9 val=l
// index=10 val=d
// index=11 val=!
// index=12 val=北
// index=15 val=京
