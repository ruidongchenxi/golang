package main
import (
	"fmt"
)
//遍历切片的每个元素。通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list{
		f(v)
	}
}
func main(){
	//使用匿名函数打印切片内容
	visit([]int{1,2,3,4}, func(v int) {
		fmt.Println(v)
	})
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main0.go
1
2
3
4
使用visit()函数将整个遍历过程进行封装，当要获取遍历期间的切片值时，只需要给visit()传入一个回调函数
准备一个整形切片传入visit()函数作为遍历的数据
定义一个匿名函数，作用打印遍历的值
*/