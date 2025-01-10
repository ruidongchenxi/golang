package main
import (
	"fmt"
)
func main(){
	for x := 0; x < 10; x++{
		for y := 0; y < 10; y++{
			if y == 2 {
				//跳转到标签
				goto breakHere
			}
			fmt.Println(y)
		}
	}
	//手动返回避免执行进入标签
	return 
	//标签
breakHere:
	fmt.Println("done")
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\流程控制\使用goto退出多层循环\main1.go
0
1
done
使用goto 语句跳转到指明的标签处，标签在18行
标签只能被goto使用，但不影响代码执行流程，此处如果不手动返回，在不满足条件时，也会执行标签里代码
定义标签
使用goto 语句后，无须额外变量就可以快速退出循环
*/