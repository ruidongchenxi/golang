package main
import (
	"fmt"
)
func main() {
	OuterLoop:
		for i := 0; i < 2; i++{
			for j := 0; j < 5; j++{
				switch j {
				case 2:
					fmt.Println(i,j)
					break OuterLoop
				case 3:
					fmt.Println(i,j)
					break OuterLoop
				}
			}
		}
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\流程控制\跳出指定循环\main0.go
0 2
外层循环标签
双层循环
使用switch进行数值分支判断
退出OuterLoop对应的循环之外，也就是
*/