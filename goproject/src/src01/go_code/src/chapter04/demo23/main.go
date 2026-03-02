package main
import (
	"fmt"
	//"strconv"
	//"math/rand"
	//"time"
)
func main() {
	label1: //标签1
	for i:=0;i<4;i++{
		label2: //标签2
		for j:=0; j< 10;j++{
			fmt.Println("j=",j)
			if j == 2{
				break label1 // 终止标签1的循环
			}else{
				break label2 // 终止标签2的循环
			}
				
		}
	}

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo21\main.go
// j= 0
// j= 0
// j= 0
// j= 0