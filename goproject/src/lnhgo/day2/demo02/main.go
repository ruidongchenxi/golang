package main
import (
	"fmt"
)
func main(){
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day2\demo02\main.go
// 0-0
// 0-1
// 结束for循环