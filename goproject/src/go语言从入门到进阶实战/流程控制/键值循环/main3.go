package main
import (
	"fmt"
)
func main(){
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\流程控制\键值循环\main3.go
1
2
3
4
创建了一个整型类型的通道
启动一个goroutine 。其逻辑的实现体现在8~11行，实现管道功能是往管道中推送数据1、2、3、4、然后关闭管道
这段goroutine在声明结束后，在第12行马上被并行执行
14行使用for range对通道c进行遍历，其实就是不断地从通道中取数据，直到关闭通道
*/