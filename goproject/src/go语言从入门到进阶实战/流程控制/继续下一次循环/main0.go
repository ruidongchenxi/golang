package main
import (
	"fmt"
)
func main(){
	OuterLoop:
		for i := 0; i < 2; i++{
			for j :=0; j < 5; j++ {
				switch j {
				case 2:
					fmt.Println(i,j)
					continue OuterLoop
				}
			}
		}
}
/*
结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\流程控制\继续下一次循环\main0.go  
0 2
1 2
*/