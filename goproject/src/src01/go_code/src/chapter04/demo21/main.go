package main
import (
	"fmt"

)
func main() {
	
	for j:=0; j< 5;j++{
			
		if j == 2{
			continue
		}
			fmt.Println("j=",j)		
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo21\main.go
// j= 0
// j= 1
// j= 3
// j= 4
