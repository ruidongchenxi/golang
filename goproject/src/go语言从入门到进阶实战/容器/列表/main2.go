package main
import (
	"fmt"
	"container/list"
)
func main(){
	l := list.New()
	l.PushBack("canon")
	l.PushFront("67")
	for i := l.Front(); i != nil; i = i.Next(){
		fmt.Println(i.Value)
	}
}
/*
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\列表\main2.go
67
canon
创建一个实例
将Canon放入列表尾部
在队列头部放入67
使用for语句进行遍历，其中i:=l.Front()表示初始赋值。只会在一开始执行一次；每次循环会进行一次i!=nil语句判断，如果返回false，表示退出循环，反之则会执行i=i.Next()
使用遍历返回的*list.Element的value成员取得放入列表时的原值
*/